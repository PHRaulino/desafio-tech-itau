package usecases

import (
	"context"
	"errors"

	"github.com/google/uuid"
	ingressosCore "github.com/phraulino/cinetuber/pkgs/ingressos/core"
	ingressosUseCase "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	pedidosUseCase "github.com/phraulino/cinetuber/pkgs/pedidos/usecases"
	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
)

type CriaReservaUseCase interface {
	Execute(ctx context.Context, sessaoID, usuarioID, assentoID, tipoIngresso string) error
}

type CriaReservaUseCaseImpl struct {
	repoReserva            core.RepoReserva
	criaPedido             pedidosUseCase.CriaPedidoUseCase
	criaIngresso           ingressosUseCase.CriaIngressoUseCase
	adicionaIngressoPedido pedidosUseCase.AdicionaItemPedidoUseCase
	atualizaIngresso       ingressosUseCase.AtualizaIngressoUseCase
	buscaIngressos         ingressosUseCase.BuscaIngressoUseCase
}

func NewCriaReservaUseCase(
	repoReserva core.RepoReserva,
	criaPedido pedidosUseCase.CriaPedidoUseCase,
	criaIngresso ingressosUseCase.CriaIngressoUseCase,
	adicionaIngressoPedido pedidosUseCase.AdicionaItemPedidoUseCase,
	atualizaIngresso ingressosUseCase.AtualizaIngressoUseCase,
	buscaIngressos ingressosUseCase.BuscaIngressoUseCase,
) CriaReservaUseCase {
	return &CriaReservaUseCaseImpl{
		repoReserva:            repoReserva,
		criaPedido:             criaPedido,
		criaIngresso:           criaIngresso,
		adicionaIngressoPedido: adicionaIngressoPedido,
		atualizaIngresso:       atualizaIngresso,
		buscaIngressos:         buscaIngressos,
	}
}

func (c *CriaReservaUseCaseImpl) Execute(ctx context.Context, sessaoID, usuarioID, assentoID, tipoIngresso string) error {
	var err error
	var pedidoID string
	var ingressosSessaoAssento []*ingressosCore.Ingresso
	var ingressoID string

	reservaDisponivel := true

	ingressosSessaoAssento, err = c.buscaIngressos.Execute(ctx, ingressosCore.BuscaIngresso{
		SessaoID:  &sessaoID,
		AssentoID: &assentoID,
	})

	for _, ingresso := range ingressosSessaoAssento {
		if (ingresso.Status == "reservado") || (ingresso.Status == "confirmado") {
			reservaDisponivel = false
		}

		if ingresso.UsuarioID == usuarioID {
			ingressoID = ingresso.IngressoID
		}
	}

	if !reservaDisponivel && ingressoID != "" {
		return nil
	}

	if !reservaDisponivel {
		return errors.New("assento não disponível para reserva")
	}

	if ingressoID != "" {
		err = c.atualizaIngresso.Execute(ctx, ingressoID, "reservado")
		if err != nil {
			return err
		}
		err = c.repoReserva.ReservaAssento(ctx, ingressoID)
		if err != nil {
			return err
		}
		return nil
	}

	ingressoID = uuid.New().String()

	pedidoID, err = c.criaPedido.Execute(ctx, usuarioID)
	if err != nil {
		return err
	}

	err = c.criaIngresso.Execute(ctx, ingressoID, sessaoID, usuarioID, assentoID, tipoIngresso)
	if err != nil {
		return err
	}

	err = c.repoReserva.ReservaAssento(ctx, ingressoID)
	if err != nil {
		return err
	}

	err = c.adicionaIngressoPedido.Execute(ctx, pedidoID, ingressoID, "ingresso", 1)
	if err != nil {
		return err
	}

	return nil
}

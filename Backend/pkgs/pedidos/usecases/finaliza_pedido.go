package usecases

import (
	"context"

	ingressosUseCases "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	"github.com/phraulino/cinetuber/pkgs/pedidos/errors"
)

type FinalizaPedidoUseCase interface {
	Execute(ctx context.Context, pedidoID string) error
}

type FinalizaPedidoUseCaseImpl struct {
	repo             core.RepoPedidos
	atualizaIngresso ingressosUseCases.AtualizaIngressoUseCase
}

func NewFinalizaPedidoUseCase(
	repo core.RepoPedidos,
	atualizaIngresso ingressosUseCases.AtualizaIngressoUseCase,
) FinalizaPedidoUseCase {
	return &FinalizaPedidoUseCaseImpl{
		repo:             repo,
		atualizaIngresso: atualizaIngresso,
	}
}

func (c *FinalizaPedidoUseCaseImpl) Execute(ctx context.Context, pedidoID string) error {
	pedido, err := c.repo.ConsultaPedido(ctx, pedidoID)
	if err != nil {
		return err
	}

	var IDSIngressos []string

	for _, item := range pedido.Itens {
		if item.Tipo == "ingresso" && item.DadosIngresso != nil {
			IDSIngressos = append(IDSIngressos, item.DadosIngresso.AssentoID)
		}

		if item.Status != "reservado" {
			return errors.ErrPedidoNaoPodeSerFinalizado
		}
	}

	for _, item := range IDSIngressos {
		err := c.atualizaIngresso.Execute(ctx, item, "confirmado")
		if err != nil {
			return err
		}
	}
	
	err = c.repo.AtualizaStatusPedido(ctx, pedidoID, "pago")
	if err != nil {
		return err
	}

	return nil
}

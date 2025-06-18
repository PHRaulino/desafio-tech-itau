package usecases

import (
	"context"

	ingressosUseCases "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
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

	for _, item := range pedido.Itens {
		if item.Tipo == "ingresso" && item.DadosIngresso != nil {
			err := c.atualizaIngresso.Execute(ctx, item.DadosIngresso.IngressoID, "confirmado")
			if err != nil {
				return err
			}
		}
	}

	err = c.repo.AtualizaStatusPedido(ctx, pedidoID, "pago")
	if err != nil {
		return err
	}

	return nil
}

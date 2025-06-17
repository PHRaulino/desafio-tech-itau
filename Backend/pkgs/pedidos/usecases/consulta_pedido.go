package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
)

type ConsultaPedidoUseCase interface {
	Execute(ctx context.Context, pedidoID string) (*core.Pedido, error)
}

type ConsultaPedidoUseCaseImpl struct {
	repo core.RepoPedidos
}

func NewConsultaPedidoUseCase(repo core.RepoPedidos) ConsultaPedidoUseCase {
	return &ConsultaPedidoUseCaseImpl{repo: repo}
}

func (c *ConsultaPedidoUseCaseImpl) Execute(ctx context.Context, pedidoID string) (*core.Pedido, error) {
	pedido, err := c.repo.ConsultaPedido(ctx, pedidoID)
	if err != nil {
		return nil, err
	}
	return pedido, nil
}

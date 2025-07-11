package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	"github.com/phraulino/cinetuber/pkgs/pedidos/errors"
)

type ConsultaPedidoUseCase interface {
	Execute(ctx context.Context, pedidoID string) (*core.PedidoCompleto, error)
}

type ConsultaPedidoUseCaseImpl struct {
	repo core.RepoPedidos
}

func NewConsultaPedidoUseCase(repo core.RepoPedidos) ConsultaPedidoUseCase {
	return &ConsultaPedidoUseCaseImpl{repo: repo}
}

func (c *ConsultaPedidoUseCaseImpl) Execute(ctx context.Context, pedidoID string) (*core.PedidoCompleto, error) {
	pedido, err := c.repo.ConsultaPedido(ctx, pedidoID)
	if err != nil {
		return nil, errors.ErrPedidoNaoEncontrado
	}
	return pedido, nil
}

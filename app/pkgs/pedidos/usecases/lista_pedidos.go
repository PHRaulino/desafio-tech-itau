package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	"github.com/phraulino/cinetuber/pkgs/pedidos/errors"
)

type ListaPedidosUseCase interface {
	Execute(ctx context.Context, filtros *core.PedidosFiltros) ([]*core.Pedido, error)
}

type ListaPedidosUseCaseImpl struct {
	repo core.RepoPedidos
}

func NewListaPedidosUseCase(repo core.RepoPedidos) ListaPedidosUseCase {
	return &ListaPedidosUseCaseImpl{repo: repo}
}

func (c *ListaPedidosUseCaseImpl) Execute(ctx context.Context, filtros *core.PedidosFiltros) ([]*core.Pedido, error) {
	pedidos, err := c.repo.ListaPedidos(ctx, filtros)
	if err != nil {
		return nil, errors.ErrPedidoNaoEncontrado
	}
	return pedidos, nil
}

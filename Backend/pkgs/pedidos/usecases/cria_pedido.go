package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
)

type CriaPedidoUseCase interface {
	Execute(ctx context.Context, usuarioID string) (string, error)
}

type CriaPedidoUseCaseImpl struct {
	repo core.RepoPedidos
}

func NewCriaPedidoUseCase(repo core.RepoPedidos) CriaPedidoUseCase {
	return &CriaPedidoUseCaseImpl{repo: repo}
}

func (c *CriaPedidoUseCaseImpl) Execute(ctx context.Context, usuarioID string) (string, error) {
	var err error
	var pedidoID string

	pedidoID = c.repo.BuscaPedidoPendente(ctx, usuarioID)
	if pedidoID != "" {
		return pedidoID, nil
	}

	pedidoID, err = c.repo.CriaPedido(ctx, usuarioID)
	if err != nil {
		return "", err
	}
	return pedidoID, nil
}

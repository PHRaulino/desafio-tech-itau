package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
)

type ICriaPedidoUseCase interface {
	Execute(ctx context.Context, usuarioID string) (string, error)
}


type CriaPedidoUseCase struct {
	repo core.RepoPedidos
}

func NewCriaPedidoUseCase(repo core.RepoPedidos) ICriaPedidoUseCase {
	return &CriaPedidoUseCase{repo: repo}
}

func (c *CriaPedidoUseCase) Execute(ctx context.Context, usuarioID string) (string, error) {
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

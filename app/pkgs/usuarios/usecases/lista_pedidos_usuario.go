package usecases

import (
	"context"

	pedidosCore "github.com/phraulino/cinetuber/pkgs/pedidos/core"
	pedidosUseCases "github.com/phraulino/cinetuber/pkgs/pedidos/usecases"
	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
)

type ListaPedidosUsuarioUseCase interface {
	Execute(ctx context.Context, usuarioID string) ([]*pedidosCore.Pedido, error)
}

type ListaPedidosUsuarioUseCaseImpl struct {
	repo                  core.RepoUsuarios
	buscaIngressosUseCase pedidosUseCases.ListaPedidosUseCase
}

func NewListaPedidosUsuarioUseCase(repo core.RepoUsuarios, buscaIngressosUseCase pedidosUseCases.ListaPedidosUseCase) ListaPedidosUsuarioUseCase {
	return &ListaPedidosUsuarioUseCaseImpl{repo: repo, buscaIngressosUseCase: buscaIngressosUseCase}
}

func (c *ListaPedidosUsuarioUseCaseImpl) Execute(ctx context.Context, usuarioID string) ([]*pedidosCore.Pedido, error) {
	pedidos, err := c.buscaIngressosUseCase.Execute(ctx, &pedidosCore.PedidosFiltros{
		UsuarioID: &usuarioID,
	})
	if err != nil {
		return nil, err
	}
	return pedidos, nil
}

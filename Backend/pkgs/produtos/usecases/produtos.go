package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/produtos/core"
)

type ListaProdutosUseCase interface {
	Execute(ctx context.Context) ([]*core.Produto, error)
}

type ListaProdutosUseCaseImpl struct {
	repo core.RepoProdutos
}

func NewListaProdutosUseCase(repo core.RepoProdutos) ListaProdutosUseCase {
	return &ListaProdutosUseCaseImpl{repo: repo}
}

func (c *ListaProdutosUseCaseImpl) Execute(ctx context.Context) ([]*core.Produto, error) {
	Produtos, err := c.repo.ListaProdutos(ctx)
	if err != nil {
		return nil, err
	}
	return Produtos, nil
}

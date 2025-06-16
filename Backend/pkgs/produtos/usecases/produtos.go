package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/produtos/core"
)

type ListaProdutosUseCase struct {
	repo core.RepoProdutos
}

func NewListaProdutosUseCase(repo core.RepoProdutos) *ListaProdutosUseCase {
	return &ListaProdutosUseCase{repo: repo}
}

func (c *ListaProdutosUseCase) Execute(ctx context.Context) ([]*core.Produto, error) {
	Produtos, err := c.repo.ListaProdutos(ctx)
	if err != nil {
		return nil, err
	}
	return Produtos, nil
}

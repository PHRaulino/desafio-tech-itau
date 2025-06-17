package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/produtos/core"
)

type ListaProdutosPorComboUseCase interface {
	Execute(ctx context.Context, comboID string) ([]*core.Produto, error)
}

type ListaProdutosPorComboUseCaseImpl struct {
	repo core.RepoProdutos
}

func NewListaProdutosPorComboUseCase(repo core.RepoProdutos) ListaProdutosPorComboUseCase {
	return &ListaProdutosPorComboUseCaseImpl{repo: repo}
}

func (c *ListaProdutosPorComboUseCaseImpl) Execute(ctx context.Context, comboID string) ([]*core.Produto, error) {
	Produtos, err := c.repo.ListaProdutosPorCombo(ctx, comboID)
	if err != nil {
		return nil, err
	}
	return Produtos, nil
}

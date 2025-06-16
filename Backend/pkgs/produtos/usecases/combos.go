package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/produtos/core"
)

type ListaCombosUseCase struct {
	repo core.RepoProdutos
}

func NewListaCombosUseCase(repo core.RepoProdutos) *ListaCombosUseCase {
	return &ListaCombosUseCase{repo: repo}
}

func (c *ListaCombosUseCase) Execute(ctx context.Context) ([]*core.Combo, error) {
	Combos, err := c.repo.ListaCombos(ctx)
	if err != nil {
		return nil, err
	}
	return Combos, nil
}

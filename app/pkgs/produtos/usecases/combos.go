package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/produtos/core"
)

type ListaCombosUseCase interface {
	Execute(ctx context.Context) ([]*core.Combo, error)
}

type ListaCombosUseCaseImpl struct {
	repo core.RepoProdutos
}

func NewListaCombosUseCase(repo core.RepoProdutos) ListaCombosUseCase {
	return &ListaCombosUseCaseImpl{repo: repo}
}

func (c *ListaCombosUseCaseImpl) Execute(ctx context.Context) ([]*core.Combo, error) {
	Combos, err := c.repo.ListaCombos(ctx)
	if err != nil {
		return nil, err
	}
	return Combos, nil
}

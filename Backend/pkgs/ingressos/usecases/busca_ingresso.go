package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
)

type BuscaIngressoUseCase interface {
	Execute(ctx context.Context, buscaIngresso core.BuscaIngresso) ([]*core.Ingresso, error)
}

type BuscaIngressoUseCaseImpl struct {
	repo core.RepoIngresso
}

func NewBuscaIngressoUseCase(repo core.RepoIngresso) BuscaIngressoUseCase {
	return &BuscaIngressoUseCaseImpl{repo: repo}
}

func (c *BuscaIngressoUseCaseImpl) Execute(ctx context.Context, buscaIngresso core.BuscaIngresso) ([]*core.Ingresso, error) {
	ingressos, err := c.repo.BuscaIngressos(ctx, buscaIngresso)
	if err != nil {
		return nil, err
	}
	return ingressos, nil
}

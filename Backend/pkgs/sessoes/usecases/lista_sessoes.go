package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
)

type ListaSessoesUseCase interface {
	Execute(ctx context.Context, payload *core.BuscaSessao) ([]*core.Sessao, error)
}

type ListaSessoesUseCaseImpl struct {
	repo core.RepoSessoes
}

func NewListaSessoesUseCase(repo core.RepoSessoes) ListaSessoesUseCase {
	return &ListaSessoesUseCaseImpl{repo: repo}
}

func (c *ListaSessoesUseCaseImpl) Execute(ctx context.Context, payload *core.BuscaSessao) ([]*core.Sessao, error) {
	sessoes, err := c.repo.ListaSessoes(ctx, payload)
	if err != nil {
		return nil, err
	}
	return sessoes, nil
}

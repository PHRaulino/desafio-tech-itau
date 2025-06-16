package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
)

type ListaSessoesUseCase struct {
	repo core.RepoSessoes
}

func NewListaSessoesUseCase(repo core.RepoSessoes) *ListaSessoesUseCase {
	return &ListaSessoesUseCase{repo: repo}
}

func (c *ListaSessoesUseCase) Execute(ctx context.Context, payload *core.BuscaSessao) ([]*core.Sessao, error) {
	sessoes, err := c.repo.ListaSessoes(ctx, payload)
	if err != nil {
		return nil, err
	}
	return sessoes, nil
}

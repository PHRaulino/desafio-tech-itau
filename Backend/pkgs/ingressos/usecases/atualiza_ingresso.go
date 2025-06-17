package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
)

type AtualizaIngressoUseCase interface {
	Execute(ctx context.Context, ingressoID, status string) error
}

type AtualizaIngressoUseCaseImpl struct {
	repo core.RepoIngresso
}

func NewAtualizaIngressoUseCase(repo core.RepoIngresso) AtualizaIngressoUseCase {
	return &AtualizaIngressoUseCaseImpl{repo: repo}
}

func (c *AtualizaIngressoUseCaseImpl) Execute(ctx context.Context, ingressoID, status string) error {
	err := c.repo.AtualizaStatusIngresso(ctx, ingressoID, status)
	if err != nil {
		return err
	}
	return nil
}

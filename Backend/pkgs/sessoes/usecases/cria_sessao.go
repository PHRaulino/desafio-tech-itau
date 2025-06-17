package usecases

import (
	"context"
	"time"

	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
	"github.com/phraulino/cinetuber/pkgs/sessoes/errors"
)

type CriaSessaoUseCase interface {
	Execute(ctx context.Context, payload *core.CriaSessao) (string, error)
}

type CriaSessaoUseCaseImpl struct {
	repo core.RepoSessoes
}

func NewCriaSessaoUseCase(repo core.RepoSessoes) CriaSessaoUseCase {
	return &CriaSessaoUseCaseImpl{repo: repo}
}

func (c *CriaSessaoUseCaseImpl) Execute(ctx context.Context, payload *core.CriaSessao) (string, error) {
	if payload.DataSessao.Before(time.Now()) {
		return "", errors.ErrDataDaSessaoAnteriorHoje
	}

	sessaoID, err := c.repo.CriaSessao(ctx, payload)
	if err != nil {
		return "", err
	}
	return sessaoID, nil
}

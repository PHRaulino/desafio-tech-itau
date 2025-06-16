package usecases

import (
	"context"
	"time"

	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
	"github.com/phraulino/cinetuber/pkgs/sessoes/errors"
)

type CriaSessaoUseCase struct {
	repo core.RepoSessoes
}

func NewCriaSessaoUseCase(repo core.RepoSessoes) *CriaSessaoUseCase {
	return &CriaSessaoUseCase{repo: repo}
}

func (c *CriaSessaoUseCase) Execute(ctx context.Context, payload *core.CriaSessao) (string, error) {
	if payload.DataSessao.Before(time.Now()) {
		return "", errors.ErrDataDaSessaoAnteriorHoje
	}

	sessaoID, err := c.repo.CriaSessao(ctx, payload)
	if err != nil {
		return "", err
	}
	return sessaoID, nil
}

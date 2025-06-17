package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
)

type ConsultaValorIngressoUseCase interface {
	Execute(ctx context.Context, tipoIngresso string) (*core.ValorIngresso, error)
}

type ConsultaValorIngressoUseCaseImpl struct {
	repo core.RepoIngresso
}

func NewConsultaValorIngressoUseCase(repo core.RepoIngresso) ConsultaValorIngressoUseCase {
	return &ConsultaValorIngressoUseCaseImpl{repo: repo}
}

func (c *ConsultaValorIngressoUseCaseImpl) Execute(ctx context.Context, tipoIngresso string) (*core.ValorIngresso, error) {
	valorIngresso, err := c.repo.ConsultaValor(ctx, tipoIngresso)
	if err != nil {
		return nil, err
	}
	return valorIngresso, nil
}

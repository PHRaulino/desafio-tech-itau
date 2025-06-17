package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
)

type ConsultaValorIngressoUseCase struct {
	repo core.RepoIngresso
}

func NewConsultaValorIngressoUseCase(repo core.RepoIngresso) *ConsultaValorIngressoUseCase {
	return &ConsultaValorIngressoUseCase{repo: repo}
}

func (c *ConsultaValorIngressoUseCase) Execute(ctx context.Context, tipoIngresso string) (*core.ValorIngresso, error) {
	valorIngresso, err := c.repo.ConsultaValor(ctx, tipoIngresso)
	if err != nil {
		return nil, err
	}
	return valorIngresso, nil
}

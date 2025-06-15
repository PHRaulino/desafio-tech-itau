package usecases

import (
	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
)

type ConsultaValorIngressoUseCase struct {
	repo core.RepoValorIngresso
}

func NewConsultaValorIngressoUseCase(repo core.RepoValorIngresso) *ConsultaValorIngressoUseCase {
	return &ConsultaValorIngressoUseCase{repo: repo}
}

func (c *ConsultaValorIngressoUseCase) Execute(tipoIngresso string) (*core.ValorIngresso, error) {
	valorIngresso, err := c.repo.ConsultaValor(tipoIngresso)
	if err != nil {
		return nil, err
	}
	return valorIngresso, nil
}

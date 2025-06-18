package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
)

type ListaAssentosUseCase interface {
	Execute(ctx context.Context, sessaoID string) ([]*core.SessaoAssento, error)
}

type ListaAssentosUseCaseImpl struct {
	repoSessao  core.RepoSessoes
	repoReserva core.RepoReserva
}

func NewListaAssentosUseCase(
	repoSessao core.RepoSessoes,
	repoReserva core.RepoReserva,
) ListaAssentosUseCase {
	return &ListaAssentosUseCaseImpl{
		repoSessao:  repoSessao,
		repoReserva: repoReserva,
	}
}

func (c *ListaAssentosUseCaseImpl) Execute(ctx context.Context, sessaoID string) ([]*core.SessaoAssento, error) {
	ingressosReservados, err := c.repoSessao.ListaAssentosReservados(ctx, sessaoID)
	if err != nil {
		return nil, err
	}

	for _, ingresso := range ingressosReservados {
		estaReservado, err := c.repoReserva.VerficaReserva(ctx, ingresso)
		if err != nil {
			return nil, err
		}

		if !estaReservado {
			err := c.repoSessao.LiberarAssento(ctx, ingresso)
			if err != nil {
				return nil, err
			}
		}
	}

	assentos, err := c.repoSessao.ListaAssentos(ctx, sessaoID)
	if err != nil {
		return nil, err
	}

	return assentos, nil
}

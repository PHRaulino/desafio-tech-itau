package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
)

type CriaIngressoUseCase interface {
	Execute(ctx context.Context, ingressoID, sessaoID, usuarioID, assentoID, tipoIngresso string) error
}

type CriaIngressoUseCaseImpl struct {
	repo core.RepoIngresso
}

func NewCriaIngressoUseCase(repo core.RepoIngresso) CriaIngressoUseCase {
	return &CriaIngressoUseCaseImpl{repo: repo}
}

func (c *CriaIngressoUseCaseImpl) Execute(ctx context.Context, ingressoID, sessaoID, usuarioID, assentoID, tipoIngresso string) error {
	var err error
	var valorIngresso *core.ValorIngresso

	valorIngresso, err = c.repo.ConsultaValor(ctx, tipoIngresso)
	if err != nil {
		return err
	}

	err = c.repo.CriaIngresso(ctx, core.Ingresso{
		IngressoID: ingressoID,
		SessaoID:   sessaoID,
		UsuarioID:  usuarioID,
		AssentoID:  assentoID,
		Status:     "reservado",
		Valor:      valorIngresso.Valor,
	})
	if err != nil {
		return err
	}
	return nil
}

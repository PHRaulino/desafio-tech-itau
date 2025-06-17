package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
)

type ICriaIngressoUseCase interface {
	Execute(ctx context.Context, ingressoID, sessaoID, usuarioID, assentoID, tipoIngresso string) error
}

type CriaIngressoUseCase struct {
	repo core.RepoIngresso
}

func NewCriaIngressoUseCase(repo core.RepoIngresso) ICriaIngressoUseCase {
	return &CriaIngressoUseCase{repo: repo}
}

func (c *CriaIngressoUseCase) Execute(ctx context.Context, ingressoID, sessaoID, usuarioID, assentoID, tipoIngresso string) error {
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

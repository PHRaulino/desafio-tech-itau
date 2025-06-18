package usecases

import (
	"context"

	ingressosCore "github.com/phraulino/cinetuber/pkgs/ingressos/core"
	ingressosUseCases "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
)

type ListaIngressosUsuarioUseCase interface {
	Execute(ctx context.Context, usuarioID string) ([]*ingressosCore.Ingresso, error)
}

type ListaIngressosUsuarioUseCaseImpl struct {
	repo                  core.RepoUsuarios
	buscaIngressosUseCase ingressosUseCases.BuscaIngressoUseCase
}

func NewListaIngressosUsuarioUseCase(repo core.RepoUsuarios, buscaIngressosUseCase ingressosUseCases.BuscaIngressoUseCase) ListaIngressosUsuarioUseCase {
	return &ListaIngressosUsuarioUseCaseImpl{repo: repo, buscaIngressosUseCase: buscaIngressosUseCase}
}

func (c *ListaIngressosUsuarioUseCaseImpl) Execute(ctx context.Context, usuarioID string) ([]*ingressosCore.Ingresso, error) {
	ingressos, err := c.buscaIngressosUseCase.Execute(ctx, ingressosCore.BuscaIngresso{
		UsuarioID: &usuarioID,
	})
	if err != nil {
		return nil, err
	}
	return ingressos, nil
}

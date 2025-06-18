package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
)

type ListaUsuariosUseCase interface {
	Execute(ctx context.Context) ([]*core.Usuario, error)
}

type ListaUsuariosUseCaseImpl struct {
	repo core.RepoUsuarios
}

func NewListaUsuariosUseCase(repo core.RepoUsuarios) ListaUsuariosUseCase {
	return &ListaUsuariosUseCaseImpl{repo: repo}
}

func (c *ListaUsuariosUseCaseImpl) Execute(ctx context.Context) ([]*core.Usuario, error) {
	usuarios, err := c.repo.ListaUsuarios(ctx)
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

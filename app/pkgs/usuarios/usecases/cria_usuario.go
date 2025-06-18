package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
)

type CriaUsuarioUseCase interface {
	Execute(ctx context.Context, usuarioInfos *core.Usuario) (string, error)
}

type CriaUsuarioUseCaseImpl struct {
	repo core.RepoUsuarios
}

func NewCriaUsuarioUseCase(repo core.RepoUsuarios) CriaUsuarioUseCase {
	return &CriaUsuarioUseCaseImpl{repo: repo}
}

func (c *CriaUsuarioUseCaseImpl) Execute(ctx context.Context, usuarioInfos *core.Usuario) (string, error) {
	usuario, err := c.repo.CriaUsuario(ctx, usuarioInfos)
	if err != nil {
		return "", err
	}
	return usuario, nil
}

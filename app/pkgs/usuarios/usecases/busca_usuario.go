package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
)

type BuscaUsuarioUseCase interface {
	Execute(ctx context.Context, usuarioID string) (*core.Usuario, error)
}

type BuscaUsuarioUseCaseImpl struct {
	repo core.RepoUsuarios
}

func NewBuscaUsuarioUseCase(repo core.RepoUsuarios) BuscaUsuarioUseCase {
	return &BuscaUsuarioUseCaseImpl{repo: repo}
}

func (c *BuscaUsuarioUseCaseImpl) Execute(ctx context.Context, usuarioID string) (*core.Usuario, error) {
	usuario, err := c.repo.BuscaUsuario(ctx, usuarioID)
	if err != nil {
		return nil, err
	}
	return usuario, nil
}

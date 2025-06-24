package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/filmes/core"
)

type ListaFilmesUseCase interface {
	Execute(ctx context.Context) ([]*core.Filme, error)
}

type ListaFilmesUseCaseImpl struct {
	repo core.RepoFilmes
}

func NewListaFilmesUseCase(repo core.RepoFilmes) ListaFilmesUseCase {
	return &ListaFilmesUseCaseImpl{repo: repo}
}

func (c *ListaFilmesUseCaseImpl) Execute(ctx context.Context) ([]*core.Filme, error) {
	filmes, err := c.repo.ListarTodos(ctx)
	if err != nil {
		return nil, err
	}
	return filmes, nil
}

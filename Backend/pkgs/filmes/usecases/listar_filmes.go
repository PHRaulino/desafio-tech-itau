package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/filmes/core"
)

type ListarFilmesUseCase struct {
	repo core.RepoFilmes
}

func NewListarFilmesUseCase(repo core.RepoFilmes) *ListarFilmesUseCase {
	return &ListarFilmesUseCase{repo: repo}
}

func (c *ListarFilmesUseCase) Execute(ctx context.Context) ([]*core.Filme, error) {
	filmes, err := c.repo.ListarTodos(ctx)
	if err != nil {
		return nil, err
	}
	return filmes, nil
}

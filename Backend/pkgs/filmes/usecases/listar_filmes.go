package usecases

import (
	"github.com/phraulino/cinetuber/pkgs/filmes/core"
)

type ListarFilmesUseCase struct {
	repo core.RepoFilmes
}

func NewListarFilmesUseCase(repo core.RepoFilmes) *ListarFilmesUseCase {
	return &ListarFilmesUseCase{repo: repo}
}

func (c *ListarFilmesUseCase) Execute() ([]*core.Filme, error) {
	filmes, err := c.repo.ListarTodos()
	if err != nil {
		return nil, err
	}
	return filmes, nil
}

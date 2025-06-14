package usecases

import models "github.com/phraulino/cinetuber/pkgs/filmes/core"


func ListarFilmes(repo models.RepoFilmes) ([]*models.Filme, error) {
	filmes, err := repo.ListarTodos()
	if err != nil {
		return nil, err
	}
	return filmes, nil
}
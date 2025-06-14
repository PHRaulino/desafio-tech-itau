//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	core "github.com/phraulino/cinetuber/pkgs/filmes/core"
	listaFilmesUseCase "github.com/phraulino/cinetuber/pkgs/filmes/usecases"
	RepoFilmes "github.com/phraulino/cinetuber/shared/adapters/sqlite"

	"github.com/google/wire"
)

var filmesSet = wire.NewSet(
	RepoFilmes.NewSQLLiteRepoFilmes,
	wire.Bind(new(core.RepoFilmes), new(*RepoFilmes.SQLLiteRepoFilmes)),
	listaFilmesUseCase.NewListarFilmesUseCase,
)

func InitializeFilmesHandler(db *sql.DB) *FilmesHandler {
	wire.Build(filmesSet, NewFilmesHandler)
	return &FilmesHandler{}
}

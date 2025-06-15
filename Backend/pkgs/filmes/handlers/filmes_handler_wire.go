//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/filmes/core"
	"github.com/phraulino/cinetuber/pkgs/filmes/usecases"
	RepoSQLite "github.com/phraulino/cinetuber/shared/adapters/sqlite"

	"github.com/google/wire"
)

var filmesSet = wire.NewSet(
	RepoSQLite.NewSQLLiteRepoFilmes,
	wire.Bind(new(core.RepoFilmes), new(*RepoSQLite.SQLLiteRepoFilmes)),
	usecases.NewListarFilmesUseCase,
)

func InitializeFilmesHandler(db *sql.DB) *FilmesHandler {
	wire.Build(filmesSet, NewFilmesHandler)
	return &FilmesHandler{}
}

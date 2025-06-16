//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/filmes/adapters"
	"github.com/phraulino/cinetuber/pkgs/filmes/core"
	"github.com/phraulino/cinetuber/pkgs/filmes/usecases"

	"github.com/google/wire"
)

var filmesSet = wire.NewSet(
	adapters.NewSQLLiteRepoFilmes,
	wire.Bind(new(core.RepoFilmes), new(*adapters.SQLLiteRepoFilmes)),
	usecases.NewListarFilmesUseCase,
)

func InitializeHandler(db *sql.DB) *FilmesHandler {
	wire.Build(filmesSet, NewFilmesHandler)
	return &FilmesHandler{}
}

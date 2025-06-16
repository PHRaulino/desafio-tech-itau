//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/sessoes/adapters"
	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
	"github.com/phraulino/cinetuber/pkgs/sessoes/usecases"

	"github.com/google/wire"
)

var SessoesSet = wire.NewSet(
	adapters.NewSQLLiteRepoSessoes,
	adapters.NewCacheEmMemoriaRepoReservas,
	wire.Bind(new(core.RepoSessoes), new(*adapters.SQLLiteRepoSessoes)),
	wire.Bind(new(core.RepoReserva), new(*adapters.CacheEmMemoriaRepoReservas)),
	usecases.NewCriaSessaoUseCase,
	usecases.NewListaSessoesUseCase,
	usecases.NewListaAssentosUseCase,
)

func InitializeHandler(db *sql.DB) *SessoesHandler {
	wire.Build(SessoesSet, NewSessoesHandler)
	return &SessoesHandler{}
}

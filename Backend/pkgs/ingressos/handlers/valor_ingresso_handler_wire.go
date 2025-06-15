//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
	"github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	RepoSQLite "github.com/phraulino/cinetuber/shared/adapters/sqlite"

	"github.com/google/wire"
)

var valorIngressoSet = wire.NewSet(
	RepoSQLite.NewSQLLiteRepoValorIngresso,
	wire.Bind(new(core.RepoValorIngresso), new(*RepoSQLite.SQLLiteRepoValorIngresso)),
	usecases.NewConsultaValorIngressoUseCase,
)

func InitializeConsultaValorIngressoHandler(db *sql.DB) *ConsultaValorIngressoHandler {
	wire.Build(valorIngressoSet, NewConsultaValorIngressoHandler)
	return &ConsultaValorIngressoHandler{}
}

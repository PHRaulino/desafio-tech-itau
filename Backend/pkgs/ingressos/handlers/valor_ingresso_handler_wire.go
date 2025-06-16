//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/ingressos/adapters"
	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
	"github.com/phraulino/cinetuber/pkgs/ingressos/usecases"

	"github.com/google/wire"
)

var valorIngressoSet = wire.NewSet(
	adapters.NewSQLLiteRepoValorIngresso,
	wire.Bind(new(core.RepoValorIngresso), new(*adapters.SQLLiteRepoValorIngresso)),
	usecases.NewConsultaValorIngressoUseCase,
)

func InitializeHandler(db *sql.DB) *ConsultaValorIngressoHandler {
	wire.Build(valorIngressoSet, NewConsultaValorIngressoHandler)
	return &ConsultaValorIngressoHandler{}
}

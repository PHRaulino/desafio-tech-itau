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
	adapters.NewSQLLiteRepoIngresso,
	wire.Bind(new(core.RepoIngresso), new(*adapters.SQLLiteRepoIngressos)),
	usecases.NewConsultaValorIngressoUseCase,
)

func InitializeHandler(db *sql.DB) *IngressoHandler {
	wire.Build(valorIngressoSet, NewIngressoHandler)
	return &IngressoHandler{}
}

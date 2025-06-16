//go:build wireinject
// +build wireinject

package handlers

import (
	"github.com/phraulino/cinetuber/pkgs/pagamentos/adapters"
	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
	"github.com/phraulino/cinetuber/pkgs/pagamentos/usecases"

	"github.com/google/wire"
)

var PagamentoSet = wire.NewSet(
	adapters.NewRepoPagamento,
	wire.Bind(new(core.RepoPagamento), new(*adapters.RepoPagamento)),
	usecases.NewPagamentoUseCase,
)

func InitializeHandler() *PagamentoHandler {
	wire.Build(PagamentoSet, NewPagamentoHandler)
	return &PagamentoHandler{}
}

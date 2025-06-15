//go:build wireinject
// +build wireinject

package handlers

import (
	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
	"github.com/phraulino/cinetuber/pkgs/pagamentos/usecases"
	adapters "github.com/phraulino/cinetuber/shared/adapters/pagamento"

	"github.com/google/wire"
)

var PagamentoSet = wire.NewSet(
	adapters.NewRepoPagamento,
	wire.Bind(new(core.RepoPagamento), new(*adapters.RepoPagamento)),
	usecases.NewPagamentoUseCase,
)

func InitializePagamentoHandler() *PagamentoHandler {
	wire.Build(PagamentoSet, NewPagamentoHandler)
	return &PagamentoHandler{}
}

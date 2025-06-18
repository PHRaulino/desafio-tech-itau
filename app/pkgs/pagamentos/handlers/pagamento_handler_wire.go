//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	ingressosAdapters "github.com/phraulino/cinetuber/pkgs/ingressos/adapters"
	ingressosCore "github.com/phraulino/cinetuber/pkgs/ingressos/core"
	ingressosUseCases "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	"github.com/phraulino/cinetuber/pkgs/pagamentos/adapters"
	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
	"github.com/phraulino/cinetuber/pkgs/pagamentos/usecases"
	pedidosAdapters "github.com/phraulino/cinetuber/pkgs/pedidos/adapters"
	pedidosCore "github.com/phraulino/cinetuber/pkgs/pedidos/core"
	pedidosUseCases "github.com/phraulino/cinetuber/pkgs/pedidos/usecases"

	"github.com/google/wire"
)

var PagamentoSet = wire.NewSet(
	adapters.NewRepoPagamento,
	pedidosAdapters.NewSQLLiteRepoPedidos,
	ingressosAdapters.NewSQLLiteRepoIngresso,
	wire.Bind(new(core.RepoPagamento), new(*adapters.RepoPagamento)),
	wire.Bind(new(pedidosCore.RepoPedidos), new(*pedidosAdapters.SQLLiteRepoPedidos)),
	wire.Bind(new(ingressosCore.RepoIngresso), new(*ingressosAdapters.SQLLiteRepoIngressos)),
	usecases.NewPagamentoUseCase,
	pedidosUseCases.NewFinalizaPedidoUseCase,
	pedidosUseCases.NewConsultaPedidoUseCase,
	pedidosUseCases.NewReverteCheckoutPedidoUseCase,
	ingressosUseCases.NewAtualizaIngressoUseCase,
)

func InitializeHandler(db *sql.DB) *PagamentoHandler {
	wire.Build(PagamentoSet, NewPagamentoHandler)
	return &PagamentoHandler{}
}

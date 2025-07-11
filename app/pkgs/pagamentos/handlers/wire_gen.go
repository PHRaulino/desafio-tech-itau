// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package handlers

import (
	"database/sql"
	"github.com/google/wire"
	adapters2 "github.com/phraulino/cinetuber/pkgs/ingressos/adapters"
	core3 "github.com/phraulino/cinetuber/pkgs/ingressos/core"
	"github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	adapters3 "github.com/phraulino/cinetuber/pkgs/pagamentos/adapters"
	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
	usecases3 "github.com/phraulino/cinetuber/pkgs/pagamentos/usecases"
	"github.com/phraulino/cinetuber/pkgs/pedidos/adapters"
	core2 "github.com/phraulino/cinetuber/pkgs/pedidos/core"
	usecases2 "github.com/phraulino/cinetuber/pkgs/pedidos/usecases"
)

// Injectors from pagamento_handler_wire.go:

func InitializeHandler(db *sql.DB) *PagamentoHandler {
	sqlLiteRepoPedidos := adapters.NewSQLLiteRepoPedidos(db)
	sqlLiteRepoIngressos := adapters2.NewSQLLiteRepoIngresso(db)
	atualizaIngressoUseCase := usecases.NewAtualizaIngressoUseCase(sqlLiteRepoIngressos)
	reverteCheckoutPedidoUseCase := usecases2.NewReverteCheckoutPedidoUseCase(sqlLiteRepoPedidos, atualizaIngressoUseCase)
	finalizaPedidoUseCase := usecases2.NewFinalizaPedidoUseCase(sqlLiteRepoPedidos, atualizaIngressoUseCase)
	consultaPedidoUseCase := usecases2.NewConsultaPedidoUseCase(sqlLiteRepoPedidos)
	repoPagamento := adapters3.NewRepoPagamento()
	pagamentoUseCase := usecases3.NewPagamentoUseCase(reverteCheckoutPedidoUseCase, finalizaPedidoUseCase, consultaPedidoUseCase, repoPagamento)
	pagamentoHandler := NewPagamentoHandler(pagamentoUseCase)
	return pagamentoHandler
}

// pagamento_handler_wire.go:

var PagamentoSet = wire.NewSet(adapters3.NewRepoPagamento, adapters.NewSQLLiteRepoPedidos, adapters2.NewSQLLiteRepoIngresso, wire.Bind(new(core.RepoPagamento), new(*adapters3.RepoPagamento)), wire.Bind(new(core2.RepoPedidos), new(*adapters.SQLLiteRepoPedidos)), wire.Bind(new(core3.RepoIngresso), new(*adapters2.SQLLiteRepoIngressos)), usecases3.NewPagamentoUseCase, usecases2.NewFinalizaPedidoUseCase, usecases2.NewConsultaPedidoUseCase, usecases2.NewReverteCheckoutPedidoUseCase, usecases.NewAtualizaIngressoUseCase)

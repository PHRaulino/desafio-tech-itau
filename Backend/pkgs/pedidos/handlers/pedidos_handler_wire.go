//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/pedidos/adapters"
	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	"github.com/phraulino/cinetuber/pkgs/pedidos/usecases"

	"github.com/google/wire"
)

var PedidosSet = wire.NewSet(
	adapters.NewSQLLiteRepoPedidos,
	wire.Bind(new(core.RepoPedidos), new(*adapters.SQLLiteRepoPedidos)),
	usecases.NewCriaPedidoUseCase,
	usecases.NewConsultaPedidoUseCase,
	usecases.NewAdicionaItemPedidoUseCase,
)

func InitializeHandler(db *sql.DB) *PedidosHandler {
	wire.Build(PedidosSet, NewPedidosHandler)
	return &PedidosHandler{}
}

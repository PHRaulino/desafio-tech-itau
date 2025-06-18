//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	ingressosAdapters "github.com/phraulino/cinetuber/pkgs/ingressos/adapters"
	ingressosCore "github.com/phraulino/cinetuber/pkgs/ingressos/core"
	ingressosUseCases "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	"github.com/phraulino/cinetuber/pkgs/pedidos/adapters"
	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	"github.com/phraulino/cinetuber/pkgs/pedidos/usecases"

	"github.com/google/wire"
)

var PedidosSet = wire.NewSet(
	adapters.NewSQLLiteRepoPedidos,
	ingressosAdapters.NewSQLLiteRepoIngresso,
	wire.Bind(new(core.RepoPedidos), new(*adapters.SQLLiteRepoPedidos)),
	wire.Bind(new(ingressosCore.RepoIngresso), new(*ingressosAdapters.SQLLiteRepoIngressos)),
	usecases.NewCriaPedidoUseCase,
	usecases.NewConsultaPedidoUseCase,
	usecases.NewAdicionaItemPedidoUseCase,
	usecases.NewCheckoutPedidoUseCase,
	ingressosUseCases.NewAtualizaIngressoUseCase,
)

func InitializeHandler(db *sql.DB) *PedidosHandler {
	wire.Build(PedidosSet, NewPedidosHandler)
	return &PedidosHandler{}
}

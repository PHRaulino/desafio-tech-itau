//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	ingressosAdapters "github.com/phraulino/cinetuber/pkgs/ingressos/adapters"
	ingressosCore "github.com/phraulino/cinetuber/pkgs/ingressos/core"
	ingressosUseCases "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	pedidosAdapters "github.com/phraulino/cinetuber/pkgs/pedidos/adapters"
	pedidosCore "github.com/phraulino/cinetuber/pkgs/pedidos/core"
	pedidosUseCases "github.com/phraulino/cinetuber/pkgs/pedidos/usecases"
	"github.com/phraulino/cinetuber/pkgs/sessoes/adapters"
	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
	"github.com/phraulino/cinetuber/pkgs/sessoes/usecases"

	"github.com/google/wire"
)

var SessoesSet = wire.NewSet(
	// Repos
	ingressosAdapters.NewSQLLiteRepoIngresso,
	pedidosAdapters.NewSQLLiteRepoPedidos,
	adapters.NewSQLLiteRepoSessoes,
	adapters.NewCacheEmMemoriaRepoReservas,

	// Bind repos
	wire.Bind(new(ingressosCore.RepoIngresso), new(*ingressosAdapters.SQLLiteRepoIngressos)),
	wire.Bind(new(pedidosCore.RepoPedidos), new(*pedidosAdapters.SQLLiteRepoPedidos)),
	wire.Bind(new(core.RepoSessoes), new(*adapters.SQLLiteRepoSessoes)),
	wire.Bind(new(core.RepoReserva), new(*adapters.CacheEmMemoriaRepoReservas)),

	// Use cases Sessões
	usecases.NewCriaSessaoUseCase,
	usecases.NewCriaReservaUseCase,
	usecases.NewListaSessoesUseCase,
	usecases.NewListaAssentosUseCase,

	// Use cases Ingressos
	ingressosUseCases.NewCriaIngressoUseCase,
	ingressosUseCases.NewBuscaIngressoUseCase,
	ingressosUseCases.NewAtualizaIngressoUseCase,

	// Use cases Pedidos
	pedidosUseCases.NewCriaPedidoUseCase,
	pedidosUseCases.NewAdicionaItemPedidoUseCase,
)

func InitializeHandler(db *sql.DB) *SessoesHandler {
	wire.Build(SessoesSet, NewSessoesHandler)
	return &SessoesHandler{}
}

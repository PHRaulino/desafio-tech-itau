//go:build wireinject
// +build wireinject

package handlers

import (
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/produtos/adapters"
	"github.com/phraulino/cinetuber/pkgs/produtos/core"
	"github.com/phraulino/cinetuber/pkgs/produtos/usecases"

	"github.com/google/wire"
)

var ProdutosSet = wire.NewSet(
	adapters.NewSQLLiteRepoProdutos,
	wire.Bind(new(core.RepoProdutos), new(*adapters.SQLLiteRepoProdutos)),
	usecases.NewListaCombosUseCase,
	usecases.NewListaProdutosUseCase,
	usecases.NewListaProdutosPorComboUseCase,
)

func InitializeProdutosHandler(db *sql.DB) *ProdutosHandler {
	wire.Build(ProdutosSet, NewProdutosHandler)
	return &ProdutosHandler{}
}

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
	"github.com/phraulino/cinetuber/pkgs/usuarios/adapters"
	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
	"github.com/phraulino/cinetuber/pkgs/usuarios/usecases"

	"github.com/google/wire"
)

var UsuariosSet = wire.NewSet(
	wire.Bind(new(core.RepoUsuarios), new(*adapters.SQLLiteRepoUsuarios)),
	wire.Bind(new(ingressosCore.RepoIngresso), new(*ingressosAdapters.SQLLiteRepoIngressos)),
	wire.Bind(new(pedidosCore.RepoPedidos), new(*pedidosAdapters.SQLLiteRepoPedidos)),
	adapters.NewSQLLiteRepoUsuarios,
	pedidosAdapters.NewSQLLiteRepoPedidos,
	ingressosAdapters.NewSQLLiteRepoIngresso,
	usecases.NewCriaUsuarioUseCase,
	usecases.NewListaUsuariosUseCase,
	usecases.NewBuscaUsuarioUseCase,
	usecases.NewGeraTokenUsuarioUseCase,
	usecases.NewListaIngressosUsuarioUseCase,
	usecases.NewListaPedidosUsuarioUseCase,
	pedidosUseCases.NewListaPedidosUseCase,
	ingressosUseCases.NewBuscaIngressoUseCase,
)

func InitializeHandler(db *sql.DB) *UsuariosHandler {
	wire.Build(UsuariosSet, NewUsuariosHandler)
	return &UsuariosHandler{}
}

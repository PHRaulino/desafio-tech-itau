package api

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	filmesHandler "github.com/phraulino/cinetuber/pkgs/filmes/handlers"
	ingressosHandler "github.com/phraulino/cinetuber/pkgs/ingressos/handlers"
	pagamentoHandler "github.com/phraulino/cinetuber/pkgs/pagamentos/handlers"
	pedidosHandler "github.com/phraulino/cinetuber/pkgs/pedidos/handlers"
	produtosHandler "github.com/phraulino/cinetuber/pkgs/produtos/handlers"
	sessoesHandler "github.com/phraulino/cinetuber/pkgs/sessoes/handlers"
	usuariosHandler "github.com/phraulino/cinetuber/pkgs/usuarios/handlers"
	httpAdapter "github.com/phraulino/cinetuber/shared/adapters/http/net_http"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

type Handler interface {
	RegisterRoutes(router *httpPorts.Router)
}

func New(db *sql.DB) {
	var router httpPorts.Router = httpAdapter.NewNetHTTPRouterAdapter()

	handlers := []Handler{
		filmesHandler.InitializeHandler(db),
		ingressosHandler.InitializeHandler(db),
		produtosHandler.InitializeHandler(db),
		pedidosHandler.InitializeHandler(db),
		sessoesHandler.InitializeHandler(db),
		pagamentoHandler.InitializeHandler(db),
		usuariosHandler.InitializeHandler(db),
	}

	for _, h := range handlers {
		h.RegisterRoutes(&router)
	}

	if err := router.ListenAndServe("8080"); err != nil {
		log.Fatal("Falha ao iniciar o servidor: ", err)
	}
}

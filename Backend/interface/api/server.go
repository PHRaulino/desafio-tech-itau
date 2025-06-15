package api

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	filmesHandler "github.com/phraulino/cinetuber/pkgs/filmes/handlers"
	ingressosHandler "github.com/phraulino/cinetuber/pkgs/ingressos/handlers"
	httpAdapter "github.com/phraulino/cinetuber/shared/adapters/http/net_http"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

func New() {
	db, err := sql.Open("sqlite3", "cinetuber.db")
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Erro ao fechar a conexão com o banco de dados: %v", err)
		}
	}()

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}

	var router httpPorts.Router = httpAdapter.NewNetHTTPRouterAdapter()

	filmesH := filmesHandler.InitializeFilmesHandler(db)
	filmesH.RegisterRoutes(&router)

	valorIngressosH := ingressosHandler.InitializeConsultaValorIngressoHandler(db)
	valorIngressosH.RegisterRoutes(&router)

	if err := router.ListenAndServe("8080"); err != nil {
		log.Fatal("Falha ao iniciar o servidor: ", err)
	}
}

package adapters

import (
	"context"
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/filmes/core"
	sqlcRepositorio "github.com/phraulino/cinetuber/shared/db/repositorios"
)

type SQLLiteRepoFilmes struct {
	db      *sql.DB
	queries *sqlcRepositorio.Queries
}

func NewSQLLiteRepoFilmes(db *sql.DB) *SQLLiteRepoFilmes {
	return &SQLLiteRepoFilmes{
		db:      db,
		queries: sqlcRepositorio.New(db),
	}
}

func (r *SQLLiteRepoFilmes) ListarTodos() ([]*core.Filme, error) {
	filmesSqlc, err := r.queries.ListaFilmes(context.Background())
	if err != nil {
		return nil, err
	}

	filmes := make([]*core.Filme, 0, len(filmesSqlc))
	for _, filmeSqlc := range filmesSqlc {
		filme := &core.Filme{
			Nome:          filmeSqlc.Nome,
			Descricao:     filmeSqlc.Descricao,
			Capa:          filmeSqlc.Capa,
			Lancamento:    filmeSqlc.Lancamento,
			Classificacao: filmeSqlc.Classificacao,
			Trailer:       filmeSqlc.Trailer,
		}
		filmes = append(filmes, filme)
	}

	return filmes, nil
}

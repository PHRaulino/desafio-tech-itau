package adaptadores

import (
	"context"
	"database/sql"

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

func (r *SQLLiteRepoFilmes) ListarTodos() ([]*sqlcRepositorio.Filme, error) {
	filmes, err := r.queries.ListaFilmes(context.Background())
	if err != nil {
		return nil, err
	}

	resultado := make([]*sqlcRepositorio.Filme, len(filmes))
	for i, filme := range filmes {
		resultado[i] = &filme
	}

	return resultado, nil
}

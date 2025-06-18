package adapters

import (
	"context"
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/produtos/core"
	sqlcRepositorio "github.com/phraulino/cinetuber/shared/db/repositorios"
)

type SQLLiteRepoProdutos struct {
	db      *sql.DB
	queries *sqlcRepositorio.Queries
}

func NewSQLLiteRepoProdutos(db *sql.DB) *SQLLiteRepoProdutos {
	return &SQLLiteRepoProdutos{
		db:      db,
		queries: sqlcRepositorio.New(db),
	}
}

func (r *SQLLiteRepoProdutos) ListaCombos(ctx context.Context) ([]*core.Combo, error) {
	combosSqlc, err := r.queries.ListaCombos(ctx)
	if err != nil {
		return nil, err
	}

	combos := make([]*core.Combo, 0, len(combosSqlc))

	for _, comboSqlc := range combosSqlc {
		c := &core.Combo{
			ID:        comboSqlc.ID,
			Nome:      comboSqlc.Nome,
			Valor:     comboSqlc.Valor,
			Descricao: comboSqlc.Descricao,
		}
		combos = append(combos, c)
	}

	return combos, nil
}

func (r *SQLLiteRepoProdutos) ListaProdutos(ctx context.Context) ([]*core.Produto, error) {
	produtosSqlc, err := r.queries.ListaProdutos(ctx)
	if err != nil {
		return nil, err
	}

	produtos := make([]*core.Produto, 0, len(produtosSqlc))

	for _, produtoSqlc := range produtosSqlc {
		c := &core.Produto{
			ID:        produtoSqlc.ID,
			Nome:      produtoSqlc.Nome,
			Valor:     produtoSqlc.Valor,
			Descricao: produtoSqlc.Descricao,
		}
		produtos = append(produtos, c)
	}

	return produtos, nil
}

func (r *SQLLiteRepoProdutos) ListaProdutosPorCombo(ctx context.Context, comboID string) ([]*core.Produto, error) {
	produtosSqlc, err := r.queries.ListaProdutosPorCombo(ctx, comboID)
	if err != nil {
		return nil, err
	}

	produtos := make([]*core.Produto, 0, len(produtosSqlc))

	for _, produtoSqlc := range produtosSqlc {
		c := &core.Produto{
			ID:        produtoSqlc.ID,
			Nome:      produtoSqlc.Nome,
			Valor:     produtoSqlc.Valor,
			Descricao: produtoSqlc.Descricao,
		}
		produtos = append(produtos, c)
	}

	return produtos, nil
}

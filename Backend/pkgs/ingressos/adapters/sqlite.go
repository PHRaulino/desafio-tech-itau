package adapters

import (
	"context"
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
	sqlcRepositorio "github.com/phraulino/cinetuber/shared/db/repositorios"
)

type SQLLiteRepoValorIngresso struct {
	db      *sql.DB
	queries *sqlcRepositorio.Queries
}

func NewSQLLiteRepoValorIngresso(db *sql.DB) *SQLLiteRepoValorIngresso {
	return &SQLLiteRepoValorIngresso{
		db:      db,
		queries: sqlcRepositorio.New(db),
	}
}

func (r *SQLLiteRepoValorIngresso) ConsultaValor(ctx context.Context, tipoIngresso string) (*core.ValorIngresso, error) {
	valorIngressoSqlc, err := r.queries.ConsultaValorIngresso(ctx, tipoIngresso)
	if err != nil {
		return nil, err
	}

	valorIngresso := &core.ValorIngresso{
		ID:    valorIngressoSqlc.ID,
		Tipo:  valorIngressoSqlc.Tipo,
		Valor: valorIngressoSqlc.Valor,
	}

	return valorIngresso, nil
}

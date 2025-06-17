package adapters

import (
	"context"
	"database/sql"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
	sqlcRepositorio "github.com/phraulino/cinetuber/shared/db/repositorios"
)

type SQLLiteRepoIngressos struct {
	db      *sql.DB
	queries *sqlcRepositorio.Queries
}

func NewSQLLiteRepoIngresso(db *sql.DB) *SQLLiteRepoIngressos {
	return &SQLLiteRepoIngressos{
		db:      db,
		queries: sqlcRepositorio.New(db),
	}
}

func (r *SQLLiteRepoIngressos) ConsultaValor(ctx context.Context, tipoIngresso string) (*core.ValorIngresso, error) {
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

func (r *SQLLiteRepoIngressos) AtualizaStatusIngresso(ctx context.Context, ingressoID, status string) error {
	err := r.queries.AtualizaStatusIngresso(ctx, sqlcRepositorio.AtualizaStatusIngressoParams{
		IngressoID: ingressoID,
		Status:     status,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLLiteRepoIngressos) CriaIngresso(ctx context.Context, ingresso core.Ingresso) error {
	err := r.queries.CriaIngresso(ctx, sqlcRepositorio.CriaIngressoParams{
		IngressoID: ingresso.IngressoID,
		SessaoID:   ingresso.SessaoID,
		AssentoID:  ingresso.AssentoID,
		Valor:      ingresso.Valor,
		UsuarioID:  ingresso.UsuarioID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLLiteRepoIngressos) BuscaIngressos(ctx context.Context, buscaIngresso core.BuscaIngresso) ([]*core.Ingresso, error) {
	ingressosSqlc, err := r.queries.ListaIngressos(ctx, sqlcRepositorio.ListaIngressosParams{
		SessaoID:  nil, // buscaIngresso.SessaoID,
		AssentoID: nil, // buscaIngresso.AssentoID,
		UsuarioID: nil, // buscaIngresso.UsuarioID,
	})
	if err != nil {
		return nil, err
	}

	ingressos := make([]*core.Ingresso, 0, len(ingressosSqlc))

	for _, ingressoSqlc := range ingressosSqlc {
		ingresso := &core.Ingresso{
			IngressoID: ingressoSqlc.IngressoID,
			AssentoID:  ingressoSqlc.AssentoID,
			UsuarioID:  ingressoSqlc.UsuarioID,
			SessaoID:   ingressoSqlc.SessaoID,
			Status:     ingressoSqlc.Status,
			Valor:      ingressoSqlc.Valor,
		}
		ingressos = append(ingressos, ingresso)
	}

	return ingressos, nil
}

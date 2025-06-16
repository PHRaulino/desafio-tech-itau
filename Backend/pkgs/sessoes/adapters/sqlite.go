package adapters

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
	sqlcRepositorio "github.com/phraulino/cinetuber/shared/db/repositorios"
)

type SQLLiteRepoSessoes struct {
	db      *sql.DB
	queries *sqlcRepositorio.Queries
}

func NewSQLLiteRepoSessoes(db *sql.DB) *SQLLiteRepoSessoes {
	return &SQLLiteRepoSessoes{
		db:      db,
		queries: sqlcRepositorio.New(db),
	}
}

func (r *SQLLiteRepoSessoes) CriaSessao(ctx context.Context, payload *core.CriaSessao) (string, error) {
	sessaoID := uuid.New().String()

	err := r.queries.CriaSessao(ctx, sqlcRepositorio.CriaSessaoParams{
		SessaoID:   sessaoID,
		SalaID:     payload.SalaID,
		FilmeID:    payload.FilmeID,
		DataSessao: payload.DataSessao,
	})
	if err != nil {
		return "", err
	}

	return sessaoID, nil
}

func (r *SQLLiteRepoSessoes) ListaSessoes(ctx context.Context, payload *core.BuscaSessao) ([]*core.Sessao, error) {
	sessoesSqlc, err := r.queries.ListaSessoes(ctx, sqlcRepositorio.ListaSessoesParams{
		FilmeID:    payload.FilmeID,
		SalaID:     payload.SalaID,
		CinemaID:   payload.CinemaID,
		DataSessao: payload.DataSessao,
	})
	if err != nil {
		return nil, nil
	}

	sessoes := make([]*core.Sessao, 0, len(sessoesSqlc))

	for _, sessaoSqlc := range sessoesSqlc {
		sessao := &core.Sessao{
			ID:            sessaoSqlc.ID,
			FilmeID:       sessaoSqlc.FilmeID,
			SalaID:        sessaoSqlc.SalaID,
			Cinema:        sessaoSqlc.NomeCinema,
			SalaDescricao: sessaoSqlc.SalaDescricao,
			DataSessao:    sessaoSqlc.DataSessao,
			Status:        sessaoSqlc.Status,
		}

		sessoes = append(sessoes, sessao)
	}

	return sessoes, nil
}

func (r *SQLLiteRepoSessoes) ListaAssentos(ctx context.Context, sessaoID string) ([]*core.SessaoAssento, error) {
	assentosSqlc, err := r.queries.ListaAssentos(ctx, sessaoID)
	if err != nil {
		return nil, err
	}

	assentosSessao := make([]*core.SessaoAssento, 0, len(assentosSqlc))

	for _, assentoSessaoSqlc := range assentosSqlc {
		assento := &core.SessaoAssento{
			AssentoID: assentoSessaoSqlc.AssentoID,
			SalaID:    assentoSessaoSqlc.SalaID,
			Status:    assentoSessaoSqlc.Status,
			Fileira:   assentoSessaoSqlc.Fileira,
			Numero:    assentoSessaoSqlc.Numero,
			Descricao: assentoSessaoSqlc.Descricao,
		}

		assentosSessao = append(assentosSessao, assento)
	}

	return assentosSessao, nil
}

func (r *SQLLiteRepoSessoes) ListaAssentosReservados(ctx context.Context, sessaoID string) ([]string, error) {
	assentos, err := r.queries.ListaAssentosReservados(ctx, sessaoID)
	if err != nil {
		return nil, err
	}

	return assentos, nil
}

func (r *SQLLiteRepoSessoes) LiberarAssento(ctx context.Context, ingressoID string) error {
	err := r.queries.DeletaIngresso(ctx, ingressoID)
	if err != nil {
		return err
	}

	return nil
}

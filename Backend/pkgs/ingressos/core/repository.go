package core

import "context"

type RepoIngresso interface {
	ConsultaValor(ctx context.Context, tipoIngresso string) (*ValorIngresso, error)
	CriaIngresso(ctx context.Context, ingresso Ingresso) error
	BuscaIngressos(ctx context.Context, buscaIngresso BuscaIngresso) ([]*Ingresso, error)
	AtualizaStatusIngresso(ctx context.Context, ingressoID, status string) error
}

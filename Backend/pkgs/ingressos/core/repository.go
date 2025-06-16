package core

import "context"

type RepoValorIngresso interface {
	ConsultaValor(ctx context.Context, tipoIngresso string) (*ValorIngresso, error)
}

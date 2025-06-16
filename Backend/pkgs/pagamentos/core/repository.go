package core

import "context"

type RepoPagamento interface {
	Efetuar(ctx context.Context, valor float64) (*Pagamento, error)
}

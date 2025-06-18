package core

import "context"

type RepoPagamento interface {
	Efetuar(ctx context.Context) (bool, error)
}

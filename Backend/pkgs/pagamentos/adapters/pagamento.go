package adapters

import (
	"context"
	"time"
)

type RepoPagamento struct{}

func NewRepoPagamento() *RepoPagamento {
	return &RepoPagamento{}
}

func (r *RepoPagamento) Efetuar(ctx context.Context) (bool, error) {
	time.Sleep(2 * time.Second)

	return true, nil
}

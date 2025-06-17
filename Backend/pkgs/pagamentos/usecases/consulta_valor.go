package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
)

type PagamentoUseCase interface {
	Execute(ctx context.Context, valor float64) (*core.Pagamento, error)
}

type PagamentoUseCaseImpl struct {
	repo core.RepoPagamento
}

func NewPagamentoUseCase(repo core.RepoPagamento) PagamentoUseCase {
	return &PagamentoUseCaseImpl{repo: repo}
}

func (c *PagamentoUseCaseImpl) Execute(ctx context.Context, valor float64) (*core.Pagamento, error) {
	Pagamento, err := c.repo.Efetuar(ctx, valor)
	if err != nil {
		return nil, err
	}
	return Pagamento, nil
}

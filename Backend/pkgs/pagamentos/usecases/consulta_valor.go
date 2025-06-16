package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
)

type PagamentoUseCase struct {
	repo core.RepoPagamento
}

func NewPagamentoUseCase(repo core.RepoPagamento) *PagamentoUseCase {
	return &PagamentoUseCase{repo: repo}
}

func (c *PagamentoUseCase) Execute(ctx context.Context, valor float64) (*core.Pagamento, error) {
	Pagamento, err := c.repo.Efetuar(ctx, valor)
	if err != nil {
		return nil, err
	}
	return Pagamento, nil
}

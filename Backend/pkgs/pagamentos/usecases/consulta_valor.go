package usecases

import (
	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
)

type PagamentoUseCase struct {
	repo core.RepoPagamento
}

func NewPagamentoUseCase(repo core.RepoPagamento) *PagamentoUseCase {
	return &PagamentoUseCase{repo: repo}
}

func (c *PagamentoUseCase) Execute(valor float64) (*core.Pagamento, error) {
	Pagamento, err := c.repo.Efetuar(valor)
	if err != nil {
		return nil, err
	}
	return Pagamento, nil
}

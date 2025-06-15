package adapters

import (
	"time"

	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
)

type RepoPagamento struct{}

func NewRepoPagamento() *RepoPagamento {
	return &RepoPagamento{}
}

func (r *RepoPagamento) Efetuar(valor float64) (*core.Pagamento, error) {
	time.Sleep(2 * time.Second)

	return &core.Pagamento{
		Mensagem: "Pagamento processado com sucesso",
		Valor:    valor,
	}, nil
}

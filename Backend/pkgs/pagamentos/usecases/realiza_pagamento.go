package usecases

import (
	"context"

	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
	pedidoErrors "github.com/phraulino/cinetuber/pkgs/pedidos/errors"
	pedidosUseCases "github.com/phraulino/cinetuber/pkgs/pedidos/usecases"
)

type PagamentoUseCase interface {
	Execute(ctx context.Context, pedidoID string) (*core.Pagamento, error)
}

type PagamentoUseCaseImpl struct {
	finalizaPedidoUseCase pedidosUseCases.FinalizaPedidoUseCase
	consultaPedidoUseCase pedidosUseCases.ConsultaPedidoUseCase
	repo                  core.RepoPagamento
}

func NewPagamentoUseCase(
	finalizaPedidoUseCase pedidosUseCases.FinalizaPedidoUseCase,
	consultaPedidoUseCase pedidosUseCases.ConsultaPedidoUseCase,
	repo core.RepoPagamento,
) PagamentoUseCase {
	return &PagamentoUseCaseImpl{
		repo:                  repo,
		finalizaPedidoUseCase: finalizaPedidoUseCase,
		consultaPedidoUseCase: consultaPedidoUseCase,
	}
}

func (c *PagamentoUseCaseImpl) Execute(ctx context.Context, pedidoID string) (*core.Pagamento, error) {
	pedido, err := c.consultaPedidoUseCase.Execute(ctx, pedidoID)
	if err != nil {
		return nil, pedidoErrors.ErrNenhumPedidoIndicado
	}

	if pedido.Status != "em pagamento" {
		return nil, pedidoErrors.ErrPedidoNaoEstaEmCheckout
	}

	pagamentoValido, err := c.repo.Efetuar(ctx)
	if err != nil {
		return nil, err
	}

	if pagamentoValido {
		err = c.finalizaPedidoUseCase.Execute(ctx, pedidoID)
		if err != nil {
			return nil, err
		}
	}

	return &core.Pagamento{
		Mensagem: "Pagamento realizado com sucesso!",
		Valor:    pedido.Total,
	}, nil
}

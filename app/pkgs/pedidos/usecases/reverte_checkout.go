package usecases

import (
	"context"

	ingressosUseCases "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
)

type ReverteCheckoutPedidoUseCase interface {
	Execute(ctx context.Context, pedidoID string) error
}

type ReverteCheckoutPedidoUseCaseImpl struct {
	repo             core.RepoPedidos
	atualizaIngresso ingressosUseCases.AtualizaIngressoUseCase
}

func NewReverteCheckoutPedidoUseCase(
	repo core.RepoPedidos,
	atualizaIngresso ingressosUseCases.AtualizaIngressoUseCase,
) ReverteCheckoutPedidoUseCase {
	return &ReverteCheckoutPedidoUseCaseImpl{
		repo:             repo,
		atualizaIngresso: atualizaIngresso,
	}
}

func (c *ReverteCheckoutPedidoUseCaseImpl) Execute(ctx context.Context, pedidoID string) error {
	pedido, err := c.repo.ConsultaPedido(ctx, pedidoID)
	if err != nil {
		return err
	}

	for _, item := range pedido.Itens {
		if item.Tipo == "ingresso" && item.DadosIngresso != nil {
			err := c.atualizaIngresso.Execute(ctx, item.DadosIngresso.IngressoID, "reservado")
			if err != nil {
				return err
			}
		}
	}

	err = c.repo.AtualizaStatusPedido(ctx, pedidoID, "pendente")
	if err != nil {
		return err
	}

	return nil
}

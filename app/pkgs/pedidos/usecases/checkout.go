package usecases

import (
	"context"

	ingressosUseCases "github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	"github.com/phraulino/cinetuber/pkgs/pedidos/errors"
)

type CheckoutPedidoUseCase interface {
	Execute(ctx context.Context, pedidoID string) error
}

type CheckoutPedidoUseCaseImpl struct {
	repo             core.RepoPedidos
	atualizaIngresso ingressosUseCases.AtualizaIngressoUseCase
}

func NewCheckoutPedidoUseCase(
	repo core.RepoPedidos,
	atualizaIngresso ingressosUseCases.AtualizaIngressoUseCase,
) CheckoutPedidoUseCase {
	return &CheckoutPedidoUseCaseImpl{
		repo:             repo,
		atualizaIngresso: atualizaIngresso,
	}
}

func (c *CheckoutPedidoUseCaseImpl) Execute(ctx context.Context, pedidoID string) error {
	pedido, err := c.repo.ConsultaPedido(ctx, pedidoID)
	if err != nil {
		return err
	}

	if pedido.Total <= 0 {
		return errors.ErrValorDoPedidoZerado
	}

	var IDSIngressos []string

	for _, item := range pedido.Itens {
		if item.Tipo == "ingresso" && item.DadosIngresso != nil {
			IDSIngressos = append(IDSIngressos, item.DadosIngresso.IngressoID)
		}

		if item.Status != "reservado" {
			return errors.ErrPedidoNaoPodeSerFinalizado
		}
	}

	for _, item := range IDSIngressos {
		err := c.atualizaIngresso.Execute(ctx, item, "em pagamento")
		if err != nil {
			return err
		}
	}

	err = c.repo.AtualizaStatusPedido(ctx, pedidoID, "em pagamento")
	if err != nil {
		return err
	}

	return nil
}

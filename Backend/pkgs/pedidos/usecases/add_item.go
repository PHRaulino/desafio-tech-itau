package usecases

import (
	"context"
	"errors"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
)

type AdicionaItemPedidoUseCase struct {
	repo core.RepoPedidos
}

func NewAdicionaItemPedidoUseCase(repo core.RepoPedidos) *AdicionaItemPedidoUseCase {
	return &AdicionaItemPedidoUseCase{repo: repo}
}

func (c *AdicionaItemPedidoUseCase) Execute(ctx context.Context, pedidoID string, payload core.AddItemPedido) error {
	var err error

	if payload.Tipo != "ingresso" {
		var quantidadeAtual float64
		quantidadeAtual = c.repo.VerificaQuantidadeItemPedido(ctx, pedidoID, payload.ItemID, payload.Tipo)

		if quantidadeAtual != payload.Quantidade {
			switch payload.Tipo {
			case "avulso":
				err = c.repo.RemoveProdutoPedido(ctx, pedidoID, payload.ItemID, payload.Tipo)
			case "combo":
				err = c.repo.RemoveComboPedido(ctx, pedidoID, payload.ItemID, payload.Tipo)
			}
		} else {
			return nil
		}

	}
	switch payload.Tipo {
	case "avulso":
		err = c.repo.AdicionaProdutoPedido(ctx, pedidoID, payload.ItemID, payload.Quantidade)
	case "combo":
		err = c.repo.AdicionaProdutosComboPedido(ctx, pedidoID, payload.ItemID, payload.Quantidade)
	case "ingresso":
		err = c.repo.AdicionaIngressoPedido(ctx, pedidoID, payload.ItemID)
	default:
		return errors.New("tipo de item inválido")
	}
	if err != nil {
		return err
	}
	return nil
}

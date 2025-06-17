package usecases

import (
	"context"
	"errors"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
)

type AdicionaItemPedidoUseCase interface {
	Execute(ctx context.Context, pedidoID, itemID, itemTipo string, itemQuantidade float64) error
}

type AdicionaItemPedidoUseCaseImpl struct {
	repo core.RepoPedidos
}

func NewAdicionaItemPedidoUseCase(repo core.RepoPedidos) AdicionaItemPedidoUseCase {
	return &AdicionaItemPedidoUseCaseImpl{repo: repo}
}

func (c *AdicionaItemPedidoUseCaseImpl) Execute(ctx context.Context, pedidoID, itemID, itemTipo string, itemQuantidade float64) error {
	var err error

	if itemTipo != "ingresso" {
		var quantidadeAtual float64
		quantidadeAtual = c.repo.VerificaQuantidadeItemPedido(ctx, pedidoID, itemID, itemTipo)

		if quantidadeAtual != itemQuantidade {
			switch itemTipo {
			case "avulso":
				err = c.repo.RemoveProdutoPedido(ctx, pedidoID, itemID, itemTipo)
			case "combo":
				err = c.repo.RemoveComboPedido(ctx, pedidoID, itemID, itemTipo)
			}
		} else {
			return nil
		}

	}
	switch itemTipo {
	case "avulso":
		err = c.repo.AdicionaProdutoPedido(ctx, pedidoID, itemID, itemQuantidade)
	case "combo":
		err = c.repo.AdicionaProdutosComboPedido(ctx, pedidoID, itemID, itemQuantidade)
	case "ingresso":
		err = c.repo.AdicionaIngressoPedido(ctx, pedidoID, itemID)
	default:
		return errors.New("tipo de item inválido")
	}
	if err != nil {
		return err
	}
	return nil
}

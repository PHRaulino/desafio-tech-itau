package core

import (
	"context"
)

type RepoPedidos interface {
	CriaPedido(ctx context.Context, usuarioID string) (string, error)
	BuscaPedidoPendente(ctx context.Context, usuarioID string) string
	ConsultaPedido(ctx context.Context, pedidoID string) (*PedidoCompleto, error)
	ListaPedidos(ctx context.Context, filtros *PedidosFiltros) ([]*Pedido, error)
	VerificaQuantidadeItemPedido(ctx context.Context, pedidoID, itemID, itemTipo string) float64
	AtualizaStatusPedido(ctx context.Context, pedidoID, status string) error
	AdicionaProdutoPedido(ctx context.Context, pedidoID, produtoID string, quantidade float64) error
	AdicionaProdutosComboPedido(ctx context.Context, pedidoID, comboID string, quantidade float64) error
	AdicionaIngressoPedido(ctx context.Context, pedidoID, ingressoID string) error
	RemoveProdutoPedido(ctx context.Context, pedidoID, produtoID, itemTipo string) error
	RemoveComboPedido(ctx context.Context, pedidoID, comboID, itemTipo string) error
	RemoveIngressoPedido(ctx context.Context, pedidoID, ingressoID string) error
}

package core

import "context"

type RepoProdutos interface {
	ListaCombos(ctx context.Context) ([]*Combo, error)
	ListaProdutos(ctx context.Context) ([]*Produto, error)
	ListaProdutosPorCombo(ctx context.Context, comboID string) ([]*Produto, error)
}

package core

type RepoProdutos interface {
	ListaCombos() ([]*Combo, error)
	ListaProdutos() ([]*Produto, error)
	ListaProdutosPorCombo(comboID string) ([]*Produto, error)
}

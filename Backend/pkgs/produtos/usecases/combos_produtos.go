package usecases

import (
	"github.com/phraulino/cinetuber/pkgs/produtos/core"
)

type ListaProdutosPorComboUseCase struct {
	repo core.RepoProdutos
}

func NewListaProdutosPorComboUseCase(repo core.RepoProdutos) *ListaProdutosPorComboUseCase {
	return &ListaProdutosPorComboUseCase{repo: repo}
}

func (c *ListaProdutosPorComboUseCase) Execute(comboID string) ([]*core.Produto, error) {
	Produtos, err := c.repo.ListaProdutosPorCombo(comboID)
	if err != nil {
		return nil, err
	}
	return Produtos, nil
}

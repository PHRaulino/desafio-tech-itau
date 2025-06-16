package core

import "time"

type Pedido struct {
	ID         string        `json:"id"`
	UsuarioID  string        `json:"usuario_id"`
	Status     string        `json:"status"`
	DataPedido time.Time     `json:"data"`
	Total      float64       `json:"total"`
	Itens      []*ItemPedido `json:"itens"`
}

type ItemPedido struct {
	Nome       string  `json:"nome"`
	Descricao  string  `json:"descricao"`
	Tipo       string  `json:"tipo"`
	Quantidade int64   `json:"quantidade"`
	Total      float64 `json:"total"`
}

type AddItemPedido struct {
	ItemID     string  `json:"item_id"`
	Tipo       string  `json:"tipo"`
	Quantidade float64 `json:"quantidade"`
}

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

type DadosIngresso struct {
	IngressoID string `json:"ingresso_id"`
	AssentoID string `json:"assento_id"`
	SessaoID  string `json:"sessao_id"`
}

type ItemPedido struct {
	Nome          string         `json:"nome"`
	Descricao     string         `json:"descricao"`
	Tipo          string         `json:"tipo"`
	Status        string         `json:"status"`
	DadosIngresso *DadosIngresso `json:"dados_ingresso,omitempty"`
	Quantidade    int64          `json:"quantidade"`
	Total         float64        `json:"total"`
}

type AddItemPedido struct {
	ItemID     string  `json:"item_id"`
	Tipo       string  `json:"tipo"`
	Quantidade float64 `json:"quantidade"`
}

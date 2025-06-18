package core

type Pagamento struct {
	Mensagem string  `json:"mensagem"`
	Valor    float64 `json:"valor"`
}
type InfosPagamento struct {
	PedidoID string `json:"pedido_id"`
}

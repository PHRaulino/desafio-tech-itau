package errors

import "fmt"

var (
	ErrPedidoNaoEncontrado        = fmt.Errorf("nenhum pedido encontrado")
	ErrNenhumPedidoIndicado       = fmt.Errorf("pedido não enviado")
	ErrPedidoNaoPodeSerFinalizado = fmt.Errorf("pedido não pode ser finalizado")
	ErrPedidoNaoEstaEmCheckout    = fmt.Errorf("o pedido não está em checkout")
	ErrValorDoPedidoZerado        = fmt.Errorf("o pedido não possui itens")
)

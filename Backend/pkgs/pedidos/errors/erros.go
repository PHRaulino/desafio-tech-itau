package errors

import "fmt"

var (
	ErrPedidoNaoEncontrado        = fmt.Errorf("pedido não encontrado")
	ErrNenhumPedidoIndicado       = fmt.Errorf("pedido não enviado")
	ErrPedidoNaoPodeSerFinalizado = fmt.Errorf("pedido não pode ser finalizado")
)

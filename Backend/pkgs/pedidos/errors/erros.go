package errors

import "fmt"

var (
	ErrPedidoNaoEncontrado  = fmt.Errorf("pedido não encontrado")
	ErrNenhumPedidoIndicado = fmt.Errorf("pedido não enviado")
)

package errors

import "fmt"

var ErrValorIngressoNaoEncontrado = fmt.Errorf("valor do ingresso não encontrado")
var ErrNenhumTipoEnviado = fmt.Errorf("o tipo do ingresso não foi enviado")

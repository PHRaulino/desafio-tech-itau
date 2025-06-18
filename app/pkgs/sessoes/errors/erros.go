package errors

import "fmt"

var (
	ErrSessaoNaoEncontrado        = fmt.Errorf("pedido não encontrado")
	ErrDataDaSessaoAnteriorHoje   = fmt.Errorf("a data da sessao deve ser maior do que hoje")
	ErrDataDaSessaoInvalida       = fmt.Errorf("a data da sessao não está em um formato valido, tente yyyy-mm-ddThh:mm:ss")
	ErrNenhumaSessaoValidaPassada = fmt.Errorf("nenhuma sessão valida informada")
)

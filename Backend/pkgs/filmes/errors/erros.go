// pacote para erros de filmes
package errors

import "fmt"

// ErrFilmeNaoEncontrado é o erro quando não é possível encontrar um filme específico.
var ErrNenhumFilmeEncontrado = fmt.Errorf("não há nenhum filme cadastrado")

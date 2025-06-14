// pacote para listar os filmes em cartaz
package core

type Filme struct {
	Nome string `json:"nome"`
	Descricao string `json:"descricao"`
	Capa string `json:"capa"`
	Lancamento string `json:"lancamento"`
	Classificacao string `json:"classificacao"`
	Trailer string `json:"trailer"`
}

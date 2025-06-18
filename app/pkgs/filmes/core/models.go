package core

import "time"

type Filme struct {
	ID            string    `json:"id"`
	Nome          string    `json:"nome"`
	Descricao     string    `json:"descricao"`
	Capa          string    `json:"capa"`
	Lancamento    time.Time `json:"lancamento"`
	Classificacao string    `json:"classificacao"`
	Trailer       string    `json:"trailer"`
}

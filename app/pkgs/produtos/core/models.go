package core

type Produto struct {
	ID        string  `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}
type Combo struct {
	ID        string    `json:"id"`
	Nome      string    `json:"nome"`
	Descricao string    `json:"descricao"`
	Valor     float64   `json:"valor"`
	Produtos  []Produto `json:"produtos"`
}

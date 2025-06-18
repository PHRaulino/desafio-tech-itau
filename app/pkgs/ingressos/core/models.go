package core

type ValorIngresso struct {
	ID    string  `json:"id"`
	Tipo  string  `json:"tipo"`
	Valor float64 `json:"valor"`
}

type Ingresso struct {
	IngressoID string
	SessaoID   string
	AssentoID  string
	UsuarioID  string
	Valor      float64
	Status     string
}

type BuscaIngresso struct {
	AssentoID *string `json:"assento_id"`
	SessaoID  *string `json:"sessao_id"`
	UsuarioID *string `json:"usuario_id"`
}

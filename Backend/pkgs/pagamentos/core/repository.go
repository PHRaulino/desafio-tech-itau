package core

type RepoPagamento interface {
	Efetuar(valor float64) (*Pagamento, error)
}

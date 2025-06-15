package core

type RepoValorIngresso interface {
	ConsultaValor(tipoIngresso string) (*ValorIngresso, error)
}

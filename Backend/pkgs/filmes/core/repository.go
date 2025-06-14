package core

// RepoFilmes traz a lista de filmes do banco de dados.
type RepoFilmes interface {
	ListarTodos() ([]*Filme, error)
}

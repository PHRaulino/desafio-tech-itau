package core


type RepoFilmes interface {
	ListarTodos() ([]*Filme, error)
}

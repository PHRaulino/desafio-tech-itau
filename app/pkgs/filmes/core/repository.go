package core

import "context"

type RepoFilmes interface {
	ListarTodos(ctx context.Context) ([]*Filme, error)
}

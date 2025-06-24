package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/phraulino/cinetuber/pkgs/filmes/core"
	"github.com/stretchr/testify/assert"
)

type repoFilmesMock struct {
	filmes []*core.Filme
	err    error
}

func (r *repoFilmesMock) ListarTodos(ctx context.Context) ([]*core.Filme, error) {
	return r.filmes, r.err
}

func TestListaFilmesUseCase_Execute_Sucesso(t *testing.T) {
	mockFilmes := []*core.Filme{
		{ID: "1", Nome: "Filme A"},
		{ID: "2", Nome: "Filme B"},
	}
	usecase := NewListaFilmesUseCase(&repoFilmesMock{filmes: mockFilmes})

	resultado, err := usecase.Execute(context.Background())

	assert.NoError(t, err)
	assert.Len(t, resultado, 2)
	assert.Equal(t, "Filme A", resultado[0].Nome)
	assert.Equal(t, "Filme B", resultado[1].Nome)
}

func TestListaFilmesUseCase_Execute_Erro(t *testing.T) {
	expectedErr := errors.New("erro ao buscar filmes")
	usecase := NewListaFilmesUseCase(&repoFilmesMock{err: expectedErr})

	resultado, err := usecase.Execute(context.Background())

	assert.Error(t, err)
	assert.Nil(t, resultado)
	assert.Equal(t, expectedErr, err)
}

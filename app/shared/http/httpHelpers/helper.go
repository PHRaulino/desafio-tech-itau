package httpHelpers

import (
	"context"
	"errors"

	"github.com/phraulino/cinetuber/shared/middlewares"
)

type UsuarioAuth struct {
	ID    string
	Email string
	Nome  string
}

func ComUsuarioAutenticado(ctx context.Context, usuario UsuarioAuth) context.Context {
	return context.WithValue(ctx, middlewares.UserKey, usuario)
}

func UsuarioAutenticado(ctx context.Context) (*middlewares.UsuarioInfo, error) {
	val := ctx.Value(middlewares.UserKey)
	if val == nil {
		return nil, errors.New("usuário não autenticado")
	}

	usuario, ok := val.(middlewares.UsuarioInfo)
	if !ok {
		return nil, errors.New("formato inválido para usuário autenticado")
	}

	return &usuario, nil
}

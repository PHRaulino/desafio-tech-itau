package usecases

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
	"github.com/phraulino/cinetuber/shared/config"
)

type GeraTokenUsuarioUseCase interface {
	Execute(ctx context.Context, usuarioID string) (string, error)
}

type GeraTokenUsuarioUseCaseImpl struct {
	repo core.RepoUsuarios
}

func NewGeraTokenUsuarioUseCase(repo core.RepoUsuarios) GeraTokenUsuarioUseCase {
	return &GeraTokenUsuarioUseCaseImpl{
		repo: repo,
	}
}

func (c *GeraTokenUsuarioUseCaseImpl) Execute(ctx context.Context, usuarioID string) (string, error) {
	usuario, err := c.repo.BuscaUsuario(ctx, usuarioID)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub": usuarioID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		"usuario": map[string]interface{}{
			"email": usuario.Email,
			"nome":  usuario.Nome,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.SecretJWT))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

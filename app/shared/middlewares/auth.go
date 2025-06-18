// shared/auth/middleware.go
package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/phraulino/cinetuber/shared/config"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

type UsuarioInfo struct {
	ID    string
	Email string
	Nome  string
}

type contextKey string

const UserKey contextKey = "usuario"

func Auth() func(httpPorts.HandlerFunc) httpPorts.HandlerFunc {
	return func(next httpPorts.HandlerFunc) httpPorts.HandlerFunc {
		return func(w httpPorts.Response, r httpPorts.Request) {
			authHeader := r.GetHeader("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"message":"Token ausente ou inválido"}`))
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(config.SecretJWT), nil
			})
			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"message":"Token inválido"}`))
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			usuario := UsuarioInfo{
				ID:    claims["sub"].(string),
				Email: claims["usuario"].(map[string]interface{})["email"].(string),
				Nome:  claims["usuario"].(map[string]interface{})["nome"].(string),
			}

			ctx := context.WithValue(r.Context(), UserKey, usuario)
			rComCtx := r.WithContext(ctx)

			next(w, rComCtx)
		}
	}
}

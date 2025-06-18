package adapters

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
	sqlcRepositorio "github.com/phraulino/cinetuber/shared/db/repositorios"
)

type SQLLiteRepoUsuarios struct {
	db      *sql.DB
	queries *sqlcRepositorio.Queries
}

func NewSQLLiteRepoUsuarios(db *sql.DB) *SQLLiteRepoUsuarios {
	return &SQLLiteRepoUsuarios{
		db:      db,
		queries: sqlcRepositorio.New(db),
	}
}

func (r *SQLLiteRepoUsuarios) ListaUsuarios(ctx context.Context) ([]*core.Usuario, error) {
	usuariosSqlc, err := r.queries.ListaUsuarios(ctx)
	if err != nil {
		return nil, err
	}

	usuarios := make([]*core.Usuario, 0, len(usuariosSqlc))

	for _, usuarioSqlc := range usuariosSqlc {
		c := &core.Usuario{
			ID:    usuarioSqlc.ID,
			Nome:  usuarioSqlc.Nome,
			Email: usuarioSqlc.Email,
		}
		usuarios = append(usuarios, c)
	}

	return usuarios, nil
}

func (r *SQLLiteRepoUsuarios) CriaUsuario(ctx context.Context, usuarioInfos *core.Usuario) (string, error) {
	usuarioID := uuid.New().String()

	err := r.queries.CriaUsuario(ctx, sqlcRepositorio.CriaUsuarioParams{
		ID:    usuarioID,
		Nome:  usuarioInfos.Nome,
		Email: usuarioInfos.Email,
	})
	if err != nil {
		return "", err
	}

	return usuarioID, nil
}

func (r *SQLLiteRepoUsuarios) BuscaUsuario(ctx context.Context, usuarioID string) (*core.Usuario, error) {
	usuario, err := r.queries.BuscaUsuario(ctx, usuarioID)
	if err != nil {
		return nil, err
	}

	return &core.Usuario{
		ID:    usuario.ID,
		Nome:  usuario.Nome,
		Email: usuario.Email,
	}, nil
}

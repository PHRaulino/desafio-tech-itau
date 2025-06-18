package core

import "context"

type RepoUsuarios interface {
	ListaUsuarios(ctx context.Context) ([]*Usuario, error)
	CriaUsuario(ctx context.Context, usuarioInfos *Usuario) (string, error)
	BuscaUsuario(ctx context.Context, usuario_id string) (*Usuario, error)
}

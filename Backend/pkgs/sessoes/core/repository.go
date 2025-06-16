package core

import (
	"context"
)

type RepoSessoes interface {
	CriaSessao(ctx context.Context, payload *CriaSessao) (string, error)
	ListaSessoes(ctx context.Context, payload *BuscaSessao) ([]*Sessao, error)
	ListaAssentos(ctx context.Context, sessaoID string) ([]*SessaoAssento, error)
	ListaAssentosReservados(ctx context.Context, sessaoID string) ([]string, error)
	LiberarAssento(ctx context.Context, ingressoID string) error
}

type RepoReserva interface {
	ReservaAssento(ctx context.Context, ingressoID string) error
	DeletaReserva(ctx context.Context, ingressoID string) error
	RenovaReserva(ctx context.Context, ingressoID string) error
	VerficaReserva(ctx context.Context, ingressoID string) (bool, error)
}

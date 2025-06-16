package adapters

import (
	"context"
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheEmMemoriaRepoReservas struct {
	cache *cache.Cache
}

func NewCacheEmMemoriaRepoReservas() *CacheEmMemoriaRepoReservas {
	return &CacheEmMemoriaRepoReservas{
		cache: cache.New(10*time.Minute, 15*time.Minute),
	}
}

func (r *CacheEmMemoriaRepoReservas) ReservaAssento(ctx context.Context, ingressoID string) error {
	r.cache.Set(ingressoID, true, 10*time.Minute)
	return nil
}

func (r *CacheEmMemoriaRepoReservas) DeletaReserva(ctx context.Context, ingressoID string) error {
	r.cache.Delete(ingressoID)
	return nil
}

func (r *CacheEmMemoriaRepoReservas) RenovaReserva(ctx context.Context, ingressoID string) error {
	_, found := r.cache.Get(ingressoID)
	if found {
		r.cache.Set(ingressoID, true, 10*time.Minute)
	}
	return nil
}

func (r *CacheEmMemoriaRepoReservas) VerficaReserva(ctx context.Context, ingressoID string) (bool, error) {
	_, found := r.cache.Get(ingressoID)
	return found, nil
}

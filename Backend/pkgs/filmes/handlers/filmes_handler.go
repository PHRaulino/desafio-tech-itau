package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/phraulino/cinetuber/pkgs/filmes/core"
	"github.com/phraulino/cinetuber/pkgs/filmes/usecases"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/utils"
)

type FilmesHandler struct {
	listaFilmesUseCase *usecases.ListarFilmesUseCase
}

func NewFilmesHandler(
	listaFilmesUseCase *usecases.ListarFilmesUseCase,
) *FilmesHandler {
	return &FilmesHandler{
		listaFilmesUseCase: listaFilmesUseCase,
	}
}

func (h *FilmesHandler) listarFilmes(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()
	filmes, err := h.listaFilmesUseCase.Execute(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*core.Filme `json:"data"`
	}{
		Data: filmes,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *FilmesHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	router.HandleFunc("GET /filmes", h.listarFilmes)
}

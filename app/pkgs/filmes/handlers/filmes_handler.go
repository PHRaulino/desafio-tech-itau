package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/phraulino/cinetuber/pkgs/filmes/core"
	"github.com/phraulino/cinetuber/pkgs/filmes/usecases"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/httpHelpers"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

type FilmesHandler struct {
	listaFilmesUseCase usecases.ListaFilmesUseCase
}

func NewFilmesHandler(
	listaFilmesUseCase usecases.ListaFilmesUseCase,
) *FilmesHandler {
	return &FilmesHandler{
		listaFilmesUseCase: listaFilmesUseCase,
	}
}

func (h *FilmesHandler) ListarFilmes(w httpPorts.Response, r httpPorts.Request) {
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
		httpHelpers.HTTPError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *FilmesHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	router.HandleFunc("GET /filmes", h.ListarFilmes)
}

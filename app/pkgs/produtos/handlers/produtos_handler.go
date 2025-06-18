package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/phraulino/cinetuber/pkgs/produtos/core"
	"github.com/phraulino/cinetuber/pkgs/produtos/usecases"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/utils"
)

type ProdutosHandler struct {
	listaCombosUseCase           usecases.ListaCombosUseCase
	listaProdutosUseCase         usecases.ListaProdutosUseCase
	listaProdutosPorComboUseCase usecases.ListaProdutosPorComboUseCase
}

func NewProdutosHandler(
	listaCombosUseCase usecases.ListaCombosUseCase,
	listaProdutosUseCase usecases.ListaProdutosUseCase,
	listaProdutosPorComboUseCase usecases.ListaProdutosPorComboUseCase,
) *ProdutosHandler {
	return &ProdutosHandler{
		listaCombosUseCase:           listaCombosUseCase,
		listaProdutosUseCase:         listaProdutosUseCase,
		listaProdutosPorComboUseCase: listaProdutosPorComboUseCase,
	}
}

func (h *ProdutosHandler) listaCombos(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	Combos, err := h.listaCombosUseCase.Execute(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*core.Combo `json:"data"`
	}{
		Data: Combos,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *ProdutosHandler) listaProdutos(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()
	Produtos, err := h.listaProdutosUseCase.Execute(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*core.Produto `json:"data"`
	}{
		Data: Produtos,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *ProdutosHandler) listaProdutosPorCombo(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	comboID := r.PathValue("combo_id")

	Produtos, err := h.listaProdutosPorComboUseCase.Execute(ctx, comboID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*core.Produto `json:"data"`
	}{
		Data: Produtos,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *ProdutosHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	router.HandleFunc("GET /produtos/combos", h.listaCombos)
	router.HandleFunc("GET /produtos/combos/{combo_id}", h.listaProdutosPorCombo)
	router.HandleFunc("GET /produtos", h.listaProdutos)
}

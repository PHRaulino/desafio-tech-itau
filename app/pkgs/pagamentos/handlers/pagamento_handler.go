package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/phraulino/cinetuber/pkgs/pagamentos/core"
	"github.com/phraulino/cinetuber/pkgs/pagamentos/usecases"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/httpHelpers"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
	"github.com/phraulino/cinetuber/shared/middlewares"
)

type PagamentoHandler struct {
	pagamentoUseCase usecases.PagamentoUseCase
}

func NewPagamentoHandler(
	pagamentoUseCase usecases.PagamentoUseCase,
) *PagamentoHandler {
	return &PagamentoHandler{
		pagamentoUseCase: pagamentoUseCase,
	}
}

func (h *PagamentoHandler) efetuarPagamento(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()
	_, err := httpHelpers.UsuarioAutenticado(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, "Não autorizado", http.StatusUnauthorized)
		return
	}

	bodyBytes, err := r.GetBody()
	if err != nil {
		httpHelpers.HTTPError(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	var infosPagamento core.InfosPagamento
	if err := json.Unmarshal(bodyBytes, &infosPagamento); err != nil {
		httpHelpers.HTTPError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	infos, err := h.pagamentoUseCase.Execute(ctx, infosPagamento.PedidoID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data *core.Pagamento `json:"data"`
	}{
		Data: infos,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *PagamentoHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	authMW := middlewares.Auth()
	router.HandleFunc("POST /pagamento", authMW(h.efetuarPagamento))
}

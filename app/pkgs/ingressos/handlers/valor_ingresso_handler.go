package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
	errorsIngressos "github.com/phraulino/cinetuber/pkgs/ingressos/errors"
	"github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/httpHelpers"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

type IngressoHandler struct {
	ingressoUseCase usecases.ConsultaValorIngressoUseCase
}

func NewIngressoHandler(
	ingressoUseCase usecases.ConsultaValorIngressoUseCase,
) *IngressoHandler {
	return &IngressoHandler{
		ingressoUseCase: ingressoUseCase,
	}
}

func (h *IngressoHandler) consultaValorIngresso(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	tipoIngresso := r.GetQueryParams("tipo_ingresso")

	if tipoIngresso == "" {
		httpHelpers.HTTPError(w, errorsIngressos.ErrNenhumTipoEnviado.Error(), http.StatusBadRequest)
		return
	}

	valorIngresso, err := h.ingressoUseCase.Execute(ctx, tipoIngresso)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data *core.ValorIngresso `json:"data"`
	}{
		Data: valorIngresso,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *IngressoHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	router.HandleFunc("GET /ingressos/valor", h.consultaValorIngresso)
}

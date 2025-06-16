package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/phraulino/cinetuber/pkgs/ingressos/core"
	errorsIngressos "github.com/phraulino/cinetuber/pkgs/ingressos/errors"
	"github.com/phraulino/cinetuber/pkgs/ingressos/usecases"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/utils"
)

type ConsultaValorIngressoHandler struct {
	ingressoUseCase *usecases.ConsultaValorIngressoUseCase
}

func NewConsultaValorIngressoHandler(
	ingressoUseCase *usecases.ConsultaValorIngressoUseCase,
) *ConsultaValorIngressoHandler {
	return &ConsultaValorIngressoHandler{
		ingressoUseCase: ingressoUseCase,
	}
}

func (h *ConsultaValorIngressoHandler) listarFilmes(w httpPorts.Response, r httpPorts.Request) {
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

func (h *ConsultaValorIngressoHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	router.HandleFunc("GET /ingresso/valor", h.listarFilmes)
}

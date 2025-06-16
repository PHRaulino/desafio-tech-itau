package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/phraulino/cinetuber/pkgs/sessoes/core"
	"github.com/phraulino/cinetuber/pkgs/sessoes/errors"
	"github.com/phraulino/cinetuber/pkgs/sessoes/usecases"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/utils"
)

type SessoesHandler struct {
	criaSessaoUseCase    *usecases.CriaSessaoUseCase
	listaSessoesUseCase  *usecases.ListaSessoesUseCase
	listaAssentosUseCase *usecases.ListaAssentosUseCase
}

func NewSessoesHandler(
	criaSessaoUseCase *usecases.CriaSessaoUseCase,
	listaSessoesUseCase *usecases.ListaSessoesUseCase,
	listaAssentosUseCase *usecases.ListaAssentosUseCase,
) *SessoesHandler {
	return &SessoesHandler{
		criaSessaoUseCase:    criaSessaoUseCase,
		listaSessoesUseCase:  listaSessoesUseCase,
		listaAssentosUseCase: listaAssentosUseCase,
	}
}

func (h *SessoesHandler) criaSessao(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	bodyBytes, err := r.GetBody()
	if err != nil {
		httpHelpers.HTTPError(w, "Falha ao receber body da request", http.StatusBadRequest)
		return
	}

	var payloadSessao *core.CriaSessao
	if err := json.Unmarshal(bodyBytes, &payloadSessao); err != nil {
		httpHelpers.HTTPError(w, "payload invalido para criar a sessao", http.StatusBadRequest)
		return
	}

	sessaoID, err := h.criaSessaoUseCase.Execute(ctx, payloadSessao)

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data string `json:"data"`
	}{
		Data: sessaoID,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *SessoesHandler) listaSessoes(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	var dataSessao *time.Time
	var err error

	filmeID := r.GetQueryParams("filme_id")
	salaID := r.GetQueryParams("sala_id")
	cinemaID := r.GetQueryParams("cinema_id")
	dataStr := r.GetQueryParams("data_sessao")

	// Ajuste para ponteiros nulos se vazio
	var filmeIDPtr, salaIDPtr, cinemaIDPtr *string

	if filmeID != "" {
		filmeIDPtr = &filmeID
	}
	if salaID != "" {
		salaIDPtr = &salaID
	}
	if cinemaID != "" {
		cinemaIDPtr = &cinemaID
	}

	if dataStr != "" {
		layout := "2006-01-02T15:04:05"
		parsed, err := time.Parse(layout, dataStr)
		if err != nil {
			httpHelpers.HTTPError(w, errors.ErrDataDaSessaoInvalida.Error(), http.StatusBadRequest)
			return
		}
		dataSessao = &parsed
	}

	payloadBuscaeSessoes := &core.BuscaSessao{
		FilmeID:    filmeIDPtr,
		SalaID:     salaIDPtr,
		CinemaID:   cinemaIDPtr,
		DataSessao: dataSessao,
	}

	sessoes, err := h.listaSessoesUseCase.Execute(ctx, payloadBuscaeSessoes)

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*core.Sessao `json:"data"`
	}{
		Data: sessoes,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *SessoesHandler) listaAssentos(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	sessaoID := r.PathValue("sessao_id")
	if sessaoID == "" {
		httpHelpers.HTTPError(w, errors.ErrNenhumaSessaoValidaPassada.Error(), http.StatusBadRequest)
		return
	}

	assentos, err := h.listaAssentosUseCase.Execute(ctx, sessaoID)

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*core.SessaoAssento `json:"data"`
	}{
		Data: assentos,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
	}
}

func (h *SessoesHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	router.HandleFunc("POST /sessoes", h.criaSessao)
	router.HandleFunc("GET /sessoes", h.listaSessoes)
	router.HandleFunc("GET /sessoes/{sessao_id}/assentos", h.listaAssentos)
}

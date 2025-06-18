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
	criaSessaoUseCase    usecases.CriaSessaoUseCase
	criaReservaUseCase   usecases.CriaReservaUseCase
	listaSessoesUseCase  usecases.ListaSessoesUseCase
	listaAssentosUseCase usecases.ListaAssentosUseCase
}

func NewSessoesHandler(
	criaSessaoUseCase usecases.CriaSessaoUseCase,
	criaReservaUseCase usecases.CriaReservaUseCase,
	listaSessoesUseCase usecases.ListaSessoesUseCase,
	listaAssentosUseCase usecases.ListaAssentosUseCase,
) *SessoesHandler {
	return &SessoesHandler{
		criaSessaoUseCase:    criaSessaoUseCase,
		criaReservaUseCase:   criaReservaUseCase,
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

	w.WriteHeader(http.StatusCreated)

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

func (h *SessoesHandler) reservarAssento(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	assentoID := r.PathValue("assento_id")
	if assentoID == "" {
		httpHelpers.HTTPError(w, errors.ErrNenhumaSessaoValidaPassada.Error(), http.StatusBadRequest)
		return
	}

	sessaoID := r.PathValue("sessao_id")
	if assentoID == "" {
		httpHelpers.HTTPError(w, errors.ErrNenhumaSessaoValidaPassada.Error(), http.StatusBadRequest)
		return
	}

	bodyBytes, err := r.GetBody()
	if err != nil {
		httpHelpers.HTTPError(w, "Falha ao receber body da request", http.StatusBadRequest)
		return
	}

	var payloadReserva *core.CriaReserva
	if err := json.Unmarshal(bodyBytes, &payloadReserva); err != nil {
		httpHelpers.HTTPError(w, "payload invalido para criar a reserva", http.StatusBadRequest)
		return
	}

	err = h.criaReservaUseCase.Execute(ctx, sessaoID, payloadReserva.UsuarioID, assentoID, payloadReserva.TipoIngresso)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SessoesHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	router.HandleFunc("POST /sessoes", h.criaSessao)
	router.HandleFunc("GET /sessoes", h.listaSessoes)
	router.HandleFunc("GET /sessoes/{sessao_id}/assentos", h.listaAssentos)
	router.HandleFunc("POST /sessoes/{sessao_id}/assentos/{assento_id}", h.reservarAssento)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	errosPedido "github.com/phraulino/cinetuber/pkgs/pedidos/errors"
	"github.com/phraulino/cinetuber/pkgs/pedidos/usecases"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/utils"
)

type PedidosHandler struct {
	consultaPedidoUseCase     *usecases.ConsultaPedidoUseCase
	criaPedidoUseCase         *usecases.CriaPedidoUseCase
	adicionaItemPedidoUseCase *usecases.AdicionaItemPedidoUseCase
}

func NewPedidosHandler(
	consultaPedidoUseCase *usecases.ConsultaPedidoUseCase,
	criaPedidoUseCase *usecases.CriaPedidoUseCase,
	adicionaItemPedidoUseCase *usecases.AdicionaItemPedidoUseCase,
) *PedidosHandler {
	return &PedidosHandler{
		consultaPedidoUseCase:     consultaPedidoUseCase,
		criaPedidoUseCase:         criaPedidoUseCase,
		adicionaItemPedidoUseCase: adicionaItemPedidoUseCase,
	}
}

func (h *PedidosHandler) consultaPedido(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	pedidoID := r.PathValue("pedido_id")

	if pedidoID == "" {
		httpHelpers.HTTPError(w, errosPedido.ErrNenhumPedidoIndicado.Error(), http.StatusNotFound)
		return
	}

	pedido, err := h.consultaPedidoUseCase.Execute(ctx, pedidoID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data *core.Pedido `json:"data"`
	}{
		Data: pedido,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *PedidosHandler) criaPedido(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	usuarioID := r.GetQueryParams("usuario_id")

	pedidoID, err := h.criaPedidoUseCase.Execute(ctx, usuarioID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data string `json:"data"`
	}{
		Data: pedidoID,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *PedidosHandler) adicionaItemAoPedido(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	pedidoID := r.PathValue("pedido_id")

	if pedidoID == "" {
		httpHelpers.HTTPError(w, errosPedido.ErrNenhumPedidoIndicado.Error(), http.StatusNotFound)
		return
	}

	bodyBytes, err := r.GetBody()
	if err != nil {
		httpHelpers.HTTPError(w, "Falha ao receber body da request", http.StatusBadRequest)
		return
	}

	var payloadItem []core.AddItemPedido
	if err := json.Unmarshal(bodyBytes, &payloadItem); err != nil {
		httpHelpers.HTTPError(w, "payload invalido para adicionar itens", http.StatusBadRequest)
		return
	}

	for _, item := range payloadItem {
		err := h.adicionaItemPedidoUseCase.Execute(ctx, pedidoID, item)
		if err != nil {
			httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *PedidosHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	router.HandleFunc("POST /pedidos", h.criaPedido)
	router.HandleFunc("GET /pedidos/{pedido_id}", h.consultaPedido)
	router.HandleFunc("POST /pedidos/{pedido_id}/itens", h.adicionaItemAoPedido)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/phraulino/cinetuber/pkgs/pedidos/core"
	errosPedido "github.com/phraulino/cinetuber/pkgs/pedidos/errors"
	"github.com/phraulino/cinetuber/pkgs/pedidos/usecases"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/httpHelpers"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
	"github.com/phraulino/cinetuber/shared/middlewares"
)

type PedidosHandler struct {
	consultaPedidoUseCase     usecases.ConsultaPedidoUseCase
	listaPedidoUseCase        usecases.ListaPedidosUseCase
	checkoutPedidoUseCase     usecases.CheckoutPedidoUseCase
	criaPedidoUseCase         usecases.CriaPedidoUseCase
	adicionaItemPedidoUseCase usecases.AdicionaItemPedidoUseCase
}

func NewPedidosHandler(
	consultaPedidoUseCase usecases.ConsultaPedidoUseCase,
	listaPedidoUseCase usecases.ListaPedidosUseCase,
	checkoutPedidoUseCase usecases.CheckoutPedidoUseCase,
	criaPedidoUseCase usecases.CriaPedidoUseCase,
	adicionaItemPedidoUseCase usecases.AdicionaItemPedidoUseCase,
) *PedidosHandler {
	return &PedidosHandler{
		consultaPedidoUseCase:     consultaPedidoUseCase,
		listaPedidoUseCase:        listaPedidoUseCase,
		checkoutPedidoUseCase:     checkoutPedidoUseCase,
		criaPedidoUseCase:         criaPedidoUseCase,
		adicionaItemPedidoUseCase: adicionaItemPedidoUseCase,
	}
}

func (h *PedidosHandler) consultaPedido(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	_, err := httpHelpers.UsuarioAutenticado(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, "Não autorizado", http.StatusUnauthorized)
		return
	}

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
		Data *core.PedidoCompleto `json:"data"`
	}{
		Data: pedido,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *PedidosHandler) listaPedidos(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	pedidos, err := h.listaPedidoUseCase.Execute(ctx, &core.PedidosFiltros{})
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*core.Pedido `json:"data"`
	}{
		Data: pedidos,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *PedidosHandler) criaPedido(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	usuarioToken, err := httpHelpers.UsuarioAutenticado(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, "Não autorizado", http.StatusUnauthorized)
		return
	}

	pedidoID, err := h.criaPedidoUseCase.Execute(ctx, usuarioToken.ID)
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

func (h *PedidosHandler) checkoutPedido(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	_, err := httpHelpers.UsuarioAutenticado(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, "Não autorizado", http.StatusUnauthorized)
		return
	}

	pedidoID := r.PathValue("pedido_id")

	err = h.checkoutPedidoUseCase.Execute(ctx, pedidoID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Mensagem string `json:"data"`
	}{
		Mensagem: "Checkout realizado com sucesso!",
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *PedidosHandler) adicionaItemAoPedido(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	_, err := httpHelpers.UsuarioAutenticado(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, "Não autorizado", http.StatusUnauthorized)
		return
	}

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
		err := h.adicionaItemPedidoUseCase.Execute(ctx, pedidoID, item.ItemID, item.Tipo, item.Quantidade)
		if err != nil {
			httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *PedidosHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	authMW := middlewares.Auth()
	router.HandleFunc("POST /pedidos", authMW(h.criaPedido))
	router.HandleFunc("GET /pedidos", h.listaPedidos)
	router.HandleFunc("GET /pedidos/{pedido_id}", authMW(h.consultaPedido))
	router.HandleFunc("POST /pedidos/{pedido_id}/checkout", authMW(h.checkoutPedido))
	router.HandleFunc("POST /pedidos/{pedido_id}/itens", authMW(h.adicionaItemAoPedido))
}

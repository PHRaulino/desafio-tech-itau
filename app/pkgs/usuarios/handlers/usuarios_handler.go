package handlers

import (
	"encoding/json"
	"net/http"

	ingressosCore "github.com/phraulino/cinetuber/pkgs/ingressos/core"
	pedidosCore "github.com/phraulino/cinetuber/pkgs/pedidos/core"
	"github.com/phraulino/cinetuber/pkgs/usuarios/core"
	"github.com/phraulino/cinetuber/pkgs/usuarios/usecases"
	httpHelpers "github.com/phraulino/cinetuber/shared/http/httpHelpers"
	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
	"github.com/phraulino/cinetuber/shared/middlewares"
)

type UsuariosHandler struct {
	criaUsuarioUseCase           usecases.CriaUsuarioUseCase
	geraTokenUsuarioUseCase      usecases.GeraTokenUsuarioUseCase
	listaUsuariosUseCase         usecases.ListaUsuariosUseCase
	listaIngressosUsuarioUseCase usecases.ListaIngressosUsuarioUseCase
	listaPedidosUsuarioUseCase   usecases.ListaPedidosUsuarioUseCase
	buscaUsuarioUseCase          usecases.BuscaUsuarioUseCase
}

func NewUsuariosHandler(
	criaUsuarioUseCase usecases.CriaUsuarioUseCase,
	geraTokenUsuarioUseCase usecases.GeraTokenUsuarioUseCase,
	listaUsuariosUseCase usecases.ListaUsuariosUseCase,
	listaIngressosUsuarioUseCase usecases.ListaIngressosUsuarioUseCase,
	listaPedidosUsuarioUseCase usecases.ListaPedidosUsuarioUseCase,
	buscaUsuarioUseCase usecases.BuscaUsuarioUseCase,
) *UsuariosHandler {
	return &UsuariosHandler{
		criaUsuarioUseCase:           criaUsuarioUseCase,
		geraTokenUsuarioUseCase:      geraTokenUsuarioUseCase,
		listaUsuariosUseCase:         listaUsuariosUseCase,
		listaIngressosUsuarioUseCase: listaIngressosUsuarioUseCase,
		listaPedidosUsuarioUseCase:   listaPedidosUsuarioUseCase,
		buscaUsuarioUseCase:          buscaUsuarioUseCase,
	}
}

func (h *UsuariosHandler) criaUsuario(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	bodyBytes, err := r.GetBody()
	if err != nil {
		httpHelpers.HTTPError(w, "Falha ao receber body da request", http.StatusBadRequest)
		return
	}

	var payloadUsuario *core.Usuario
	if err := json.Unmarshal(bodyBytes, &payloadUsuario); err != nil {
		httpHelpers.HTTPError(w, "payload invalido para adicionar itens", http.StatusBadRequest)
		return
	}

	usuarioID, err := h.criaUsuarioUseCase.Execute(ctx, payloadUsuario)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data string `json:"data"`
	}{
		Data: usuarioID,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *UsuariosHandler) listaUsuarios(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()
	Usuarios, err := h.listaUsuariosUseCase.Execute(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*core.Usuario `json:"data"`
	}{
		Data: Usuarios,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *UsuariosHandler) listaIngressosUsuarios(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	usuarioToken, err := httpHelpers.UsuarioAutenticado(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, "Não autorizado", http.StatusUnauthorized)
		return
	}

	ingressos, err := h.listaIngressosUsuarioUseCase.Execute(ctx, usuarioToken.ID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*ingressosCore.Ingresso `json:"data"`
	}{
		Data: ingressos,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *UsuariosHandler) listaPedidosUsuarios(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	usuarioToken, err := httpHelpers.UsuarioAutenticado(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, "Não autorizado", http.StatusUnauthorized)
		return
	}

	ingressos, err := h.listaPedidosUsuarioUseCase.Execute(ctx, usuarioToken.ID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data []*pedidosCore.Pedido `json:"data"`
	}{
		Data: ingressos,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *UsuariosHandler) buscaUsuario(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	usuarioToken, err := httpHelpers.UsuarioAutenticado(ctx)
	if err != nil {
		httpHelpers.HTTPError(w, "Não autorizado", http.StatusUnauthorized)
		return
	}

	usuario, err := h.buscaUsuarioUseCase.Execute(ctx, usuarioToken.ID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data *core.Usuario `json:"data"`
	}{
		Data: usuario,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *UsuariosHandler) geraToken(w httpPorts.Response, r httpPorts.Request) {
	ctx := r.Context()

	bodyBytes, err := r.GetBody()
	if err != nil {
		httpHelpers.HTTPError(w, "Falha ao receber body da request", http.StatusBadRequest)
		return
	}

	var payloadUsuario struct {
		UsuarioID string `json:"usuario_id"`
	}
	if err := json.Unmarshal(bodyBytes, &payloadUsuario); err != nil {
		httpHelpers.HTTPError(w, "payload invalido para adicionar itens", http.StatusBadRequest)
		return
	}

	token, err := h.geraTokenUsuarioUseCase.Execute(ctx, payloadUsuario.UsuarioID)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	response := struct {
		Data string `json:"data"`
	}{
		Data: token,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpHelpers.HTTPError(w, err.Error(), 409)
		return
	}
}

func (h *UsuariosHandler) RegisterRoutes(httpRouter *httpPorts.Router) {
	router := *httpRouter
	authMW := middlewares.Auth()
	router.HandleFunc("GET /usuarios", h.listaUsuarios)
	router.HandleFunc("POST /usuarios", h.criaUsuario)
	router.HandleFunc("POST /usuarios/token", h.geraToken)
	router.HandleFunc("GET /usuario/me", authMW(h.buscaUsuario))
	router.HandleFunc("GET /usuario/ingressos", authMW(h.listaIngressosUsuarios))
	router.HandleFunc("GET /usuario/pedidos", authMW(h.listaPedidosUsuarios))
}

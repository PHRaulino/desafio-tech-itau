// Pacote adapters fornece adaptadores de requisição e resposta HTTP para o backend.
package adapters

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

// NetHTTPRequestAdapter é um adaptador para requisições HTTP.
type NetHTTPRequestAdapter struct {
	req *http.Request
}

// GetBody retorna o corpo da requisição HTTP como um slice de bytes.
func (a *NetHTTPRequestAdapter) GetBody() ([]byte, error) {
	bodyBytes, err := io.ReadAll(a.req.Body)
	if err != nil {
		return nil, err
	}

	a.req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}

// GetMethod retorna o método HTTP da requisição.
func (a *NetHTTPRequestAdapter) GetMethod() string {
	return a.req.Method
}

// GetPath retorna o caminho da URL da requisição HTTP.
func (a *NetHTTPRequestAdapter) GetPath() string {
	return a.req.URL.Path
}

// GetQueryParams retorna o valor do parâmetro de consulta especificado da requisição HTTP.
func (a *NetHTTPRequestAdapter) GetQueryParams(key string) string {
	return a.req.URL.Query().Get(key)
}

// PathValue retorna o valor da chave especificada do caminho da URL.
func (a *NetHTTPRequestAdapter) PathValue(key string) string {
	return a.req.PathValue(key)
}

// GetHeader retorna o valor do cabeçalho especificado da requisição HTTP.
func (a *NetHTTPRequestAdapter) GetHeader(key string) string {
	return a.req.Header.Get(key)
}

// NetHTTPResponseAdapter é um adaptador para respostas HTTP.
type NetHTTPResponseAdapter struct {
	w http.ResponseWriter
}

// SetHeader define o valor do cabeçalho especificado na resposta HTTP.
func (a *NetHTTPResponseAdapter) SetHeader(key, value string) {
	a.w.Header().Set(key, value)
}

// WriteHeader escreve o código de status HTTP na resposta.
func (a *NetHTTPResponseAdapter) WriteHeader(statusCode int) {
	a.w.WriteHeader(statusCode)
}

// Write escreve os dados na resposta HTTP.
func (a *NetHTTPResponseAdapter) Write(data []byte) (int, error) {
	return a.w.Write(data)
}

// NetHTTPHandlerAdapter é um adaptador para manipulação de requisições e respostas HTTP.
type NetHTTPHandlerAdapter struct {
	Serve func(httpPorts.Response, httpPorts.Request)
}

// ServeHTTP manipula requisições e respostas HTTP.
func (h *NetHTTPHandlerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := &NetHTTPResponseAdapter{w: w}
	request := &NetHTTPRequestAdapter{req: r}
	h.Serve(response, request)
}

// NetHTTPHandlerFuncAdapter é um adaptador para manipulação de requisições HTTP usando uma função.
type NetHTTPHandlerFuncAdapter struct {
	Handler func(httpPorts.Response, httpPorts.Request)
}

// HandlerFunc manipula requisições e respostas HTTP.
func (h *NetHTTPHandlerFuncAdapter) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	response := &NetHTTPResponseAdapter{w: w}
	request := &NetHTTPRequestAdapter{req: r}
	h.Handler(response, request)
}

// NetHTTPRouterAdapter é um adaptador para roteamento HTTP.
type NetHTTPRouterAdapter struct {
	mux *http.ServeMux
}

// NewNetHTTPRouterAdapter cria uma nova instância de NetHTTPRouterAdapter.
func NewNetHTTPRouterAdapter() *NetHTTPRouterAdapter {
	return &NetHTTPRouterAdapter{
		mux: http.NewServeMux(),
	}
}

// Handle registra o manipulador para o caminho fornecido no roteador HTTP.
func (r *NetHTTPRouterAdapter) Handle(path string, handler httpPorts.Handler) {
	r.mux.Handle(path, &NetHTTPHandlerAdapter{Serve: handler.Serve})
}

// HandleFunc registra a função manipuladora para o caminho fornecido no roteador HTTP.
func (r *NetHTTPRouterAdapter) HandleFunc(
	path string,
	handler func(httpPorts.Response, httpPorts.Request),
) {
	handleFuncAdapter := &NetHTTPHandlerFuncAdapter{Handler: handler}
	r.mux.HandleFunc(path, handleFuncAdapter.HandlerFunc)
}

// GetMux retorna o ServeMux do HTTP.
func (r *NetHTTPRouterAdapter) GetMux() *http.ServeMux {
	return r.mux
}

// ListenAndServe inicia o servidor HTTP na porta especificada.
func (r *NetHTTPRouterAdapter) ListenAndServe(port string) error {
	fmt.Printf("Servidor rodando na porta %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r.mux)
}

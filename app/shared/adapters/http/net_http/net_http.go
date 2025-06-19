package adapters

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/phraulino/cinetuber/docs"

	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

type NetHTTPRequestAdapter struct {
	req *http.Request
}

func (a *NetHTTPRequestAdapter) GetBody() ([]byte, error) {
	bodyBytes, err := io.ReadAll(a.req.Body)
	if err != nil {
		return nil, err
	}
	a.req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes, nil
}

func (a *NetHTTPRequestAdapter) GetMethod() string {
	return a.req.Method
}

func (a *NetHTTPRequestAdapter) GetPath() string {
	return a.req.URL.Path
}

func (a *NetHTTPRequestAdapter) GetQueryParams(key string) string {
	return a.req.URL.Query().Get(key)
}

func (a *NetHTTPRequestAdapter) PathValue(key string) string {
	return a.req.PathValue(key)
}

func (a *NetHTTPRequestAdapter) GetHeader(key string) string {
	return a.req.Header.Get(key)
}

func (a *NetHTTPRequestAdapter) Context() context.Context {
	return a.req.Context()
}

// ✅ Novo método necessário
func (a *NetHTTPRequestAdapter) WithContext(ctx context.Context) httpPorts.Request {
	return &NetHTTPRequestAdapter{
		req: a.req.WithContext(ctx),
	}
}

type NetHTTPResponseAdapter struct {
	w http.ResponseWriter
}

func (a *NetHTTPResponseAdapter) SetHeader(key, value string) {
	a.w.Header().Set(key, value)
}

func (a *NetHTTPResponseAdapter) WriteHeader(statusCode int) {
	a.w.WriteHeader(statusCode)
}

func (a *NetHTTPResponseAdapter) Write(data []byte) (int, error) {
	return a.w.Write(data)
}

type NetHTTPHandlerAdapter struct {
	Serve func(httpPorts.Response, httpPorts.Request)
}

func (h *NetHTTPHandlerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := &NetHTTPResponseAdapter{w: w}
	request := &NetHTTPRequestAdapter{req: r}
	h.Serve(response, request)
}

type NetHTTPHandlerFuncAdapter struct {
	Handler func(httpPorts.Response, httpPorts.Request)
}

func (h *NetHTTPHandlerFuncAdapter) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	response := &NetHTTPResponseAdapter{w: w}
	request := &NetHTTPRequestAdapter{req: r}
	h.Handler(response, request)
}

type NetHTTPRouterAdapter struct {
	mux *http.ServeMux
}

func NewNetHTTPRouterAdapter() *NetHTTPRouterAdapter {
	return &NetHTTPRouterAdapter{
		mux: http.NewServeMux(),
	}
}

func (r *NetHTTPRouterAdapter) Handle(path string, handler httpPorts.Handler) {
	r.mux.Handle(path, &NetHTTPHandlerAdapter{Serve: handler.Serve})
}

func (r *NetHTTPRouterAdapter) HandleFunc(
	path string,
	handler func(httpPorts.Response, httpPorts.Request),
) {
	handleFuncAdapter := &NetHTTPHandlerFuncAdapter{Handler: handler}
	r.mux.HandleFunc(path, handleFuncAdapter.HandlerFunc)
}

func (r *NetHTTPRouterAdapter) GetMux() *http.ServeMux {
	return r.mux
}

func (r *NetHTTPRouterAdapter) ListenAndServe(port string) error {
	loggedHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		path := req.URL.Path
		method := req.Method

		// Captura status com um wrapper
		wrapped := &statusRecorder{ResponseWriter: w, status: 200}
		r.mux.ServeHTTP(wrapped, req)

		duration := time.Since(start)
		fmt.Printf("[HTTP] %s %s -> %d (%s) UA: %s\n",
			method,
			path,
			wrapped.status,
			duration,
			req.Header.Get("User-Agent"),
		)
	})

	fmt.Printf("Servidor rodando na porta %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), loggedHandler)
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (r *NetHTTPRouterAdapter) ServeOpenAPIDocs() {
	fs := http.FileServer(http.FS(docs.OpenAPIFS))
	r.mux.Handle("/docs/", http.StripPrefix("/docs/", fs))
}

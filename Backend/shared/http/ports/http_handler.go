// Pacote ports fornece interfaces para manipulação de requisições e respostas HTTP.
package ports

// ResponseAPI representa a estrutura de uma resposta da API.
type ResponseAPI struct {
	Message string `json:"message"`
}

// Request representa uma requisição HTTP.
type Request interface {
	GetBody() ([]byte, error)
	GetMethod() string
	GetPath() string
	GetQueryParams(key string) string
	PathValue(key string) string
}

// Response representa uma resposta HTTP.
type Response interface {
	SetHeader(key, value string)
	WriteHeader(statusCode int)
	Write(data []byte) (int, error)
}

// Handler representa um manipulador HTTP.
type Handler interface {
	Serve(w Response, r Request)
}

// Router representa um roteador HTTP.
type Router interface {
	Handle(path string, handler Handler)
	HandleFunc(path string, handler func(w Response, r Request))
	ListenAndServe(port string) error
}

package ports

import "context"

type ResponseAPI struct {
	Message string `json:"message"`
}

type Request interface {
	GetBody() ([]byte, error)
	GetMethod() string
	GetPath() string
	GetQueryParams(key string) string
	PathValue(key string) string
	Context() context.Context
	WithContext(ctx context.Context) Request
	GetHeader(key string) string
}

type Response interface {
	SetHeader(key, value string)
	WriteHeader(statusCode int)
	Write(data []byte) (int, error)
}

type Handler interface {
	Serve(w Response, r Request)
}
type (
	HandlerFunc func(Response, Request)
	Router      interface {
		Handle(path string, handler Handler)
		HandleFunc(path string, handler func(w Response, r Request))
		ListenAndServe(port string) error
	}
)

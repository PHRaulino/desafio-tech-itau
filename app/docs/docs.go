package docs

import "embed"

//go:embed index.html swagger.html redoc.html openapi.yaml assets/*
var OpenAPIFS embed.FS

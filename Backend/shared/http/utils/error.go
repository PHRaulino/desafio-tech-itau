package utils

import (
	"encoding/json"
	"log"

	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

func HTTPError(w httpPorts.Response, message string, code int) {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(httpPorts.ResponseAPI{
		Message: message,
	},
	); err != nil {
		log.Println("failed to write response:", err)
	}
}

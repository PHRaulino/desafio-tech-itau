// Pacote utils fornece funções utilitárias para lidar com erros HTTP.
package utils

import (
	"encoding/json"
	"log"

	httpPorts "github.com/phraulino/cinetuber/shared/http/ports"
)

// HTTPError escreve uma resposta de erro HTTP com a mensagem e o código de status fornecidos.
func HTTPError(w httpPorts.Response, message string, code int) {
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(httpPorts.ResponseAPI{
		Message: message,
	},
	); err != nil {
		// trata o erro, por exemplo, registrando-o no log
		log.Println("failed to write response:", err)
	}
}

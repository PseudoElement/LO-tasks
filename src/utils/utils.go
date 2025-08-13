package utils

import (
	"encoding/json"
	"net/http"

	"github.com/pseudoelement/lo-tasks/src/core/models/structs"
)

func SuccessResponse[T any](w http.ResponseWriter, body T, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func FailResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	var res = structs.MessageJson{Message: message}
	json.NewEncoder(w).Encode(res)
}

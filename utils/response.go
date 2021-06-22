package utils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	JSONResponse(w, status, errorResponse{Error: err.Error()})
}

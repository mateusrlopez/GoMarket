package responses

import "net/http"

type ErrorResponse struct {
	Error string `json:"error"`
}

func Error(w http.ResponseWriter, status int, err error) {
	JSON(w, status, ErrorResponse{Error: err.Error()})
}

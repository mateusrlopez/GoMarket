package responses

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

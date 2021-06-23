package middlewares

import (
	"errors"
	"net/http"

	"github.com/mateusrlopez/go-market/constants"
	"github.com/mateusrlopez/go-market/types"
	"github.com/mateusrlopez/go-market/utils"
)

func AdminMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(constants.ContextKey).(types.ContextPayload).User

		if !user.Admin {
			utils.ErrorResponse(w, http.StatusForbidden, errors.New("Access Forbidden"))
			return
		}

		h(w, r)
	}
}

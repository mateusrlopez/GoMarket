package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/mateusrlopez/go-market/responses"
	"github.com/mateusrlopez/go-market/utils"
)

type ContextKey string

func AuthorizationMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		splittedTokenHeader := strings.Split(tokenHeader, " ")

		if len(splittedTokenHeader) != 2 {
			responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		token := splittedTokenHeader[1]
		ctp, err := utils.ValidateToken(token)

		if err != nil {
			responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		const ctk ContextKey = "claims"
		ctx := context.WithValue(r.Context(), ctk, ctp)

		h(w, r.WithContext(ctx))
	}
}

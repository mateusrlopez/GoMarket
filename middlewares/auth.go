package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/mateusrlopez/go-market/constants"
	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/repositories"
	"github.com/mateusrlopez/go-market/responses"
	"github.com/mateusrlopez/go-market/settings"
	"github.com/mateusrlopez/go-market/types"
)

type AuthorizationMiddleware struct {
	TokenRepository repositories.TokenRepository
	UserRepository  repositories.UserRepository
}

func (m *AuthorizationMiddleware) AccessMiddleware(h http.HandlerFunc) http.HandlerFunc {
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
		tmd, err := m.TokenRepository.ValidateToken(token, settings.Settings.Server.AccessSecret)

		if err != nil {
			responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		err = m.TokenRepository.RetrieveTokenMetadata(tmd.UUID)

		if err != nil {
			responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		user := &models.User{}
		err = m.UserRepository.RetriveByID(uint(tmd.UserId), user)

		if err != nil {
			responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		ctp := types.ContextPayload{User: user, TokenId: tmd.UUID}
		ctx := context.WithValue(r.Context(), constants.ContextKey, ctp)

		h(w, r.WithContext(ctx))
	}
}

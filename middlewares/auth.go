package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/mateusrlopez/go-market/constants"
	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/repositories"
	"github.com/mateusrlopez/go-market/settings"
	"github.com/mateusrlopez/go-market/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorizationMiddleware struct {
	TokenRepository repositories.TokenRepository
	UserRepository  repositories.UserRepository
}

func (m *AuthorizationMiddleware) AccessMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		splittedTokenHeader := strings.Split(tokenHeader, " ")

		if len(splittedTokenHeader) != 2 {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		token := splittedTokenHeader[1]
		tmd, err := m.TokenRepository.ValidateToken(token, settings.Settings.Server.AccessSecret)

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		if result := m.TokenRepository.RetrieveTokenMetadata(tmd.UserId, "access"); result.Err() != nil || result.Val() != tmd.UUID {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		user := &models.User{}
		id, err := primitive.ObjectIDFromHex(tmd.UserId)

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		err = m.UserRepository.RetriveByID(id, user)

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		ctx := context.WithValue(r.Context(), constants.ContextKey, user)

		h(w, r.WithContext(ctx))
	}
}

func (m *AuthorizationMiddleware) RefreshMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		splittedTokenHeader := strings.Split(tokenHeader, " ")

		if len(splittedTokenHeader) != 2 {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		token := splittedTokenHeader[1]
		tmd, err := m.TokenRepository.ValidateToken(token, settings.Settings.Server.RefreshSecret)

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		if result := m.TokenRepository.RetrieveTokenMetadata(tmd.UserId, "refresh"); result.Err() != nil || result.Val() != tmd.UUID {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		user := &models.User{}
		id, err := primitive.ObjectIDFromHex(tmd.UserId)

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		err = m.UserRepository.RetriveByID(id, user)

		if err != nil {
			utils.ErrorResponse(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}

		ctx := context.WithValue(r.Context(), constants.ContextKey, user)

		h(w, r.WithContext(ctx))
	}
}

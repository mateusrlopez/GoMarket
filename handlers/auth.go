package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mateusrlopez/go-market/constants"
	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/repositories"
	"github.com/mateusrlopez/go-market/responses"
	"github.com/mateusrlopez/go-market/types"
)

type AuthHandler struct {
	UserRepository  repositories.UserRepository
	TokenRepository repositories.TokenRepository
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.ValidateRegister()

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = h.UserRepository.Create(&user)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	tr, err := h.TokenRepository.GenerateTokens(user.ID)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, tr)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.ValidateLogin()

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	retrievedUser := models.User{}
	err = h.UserRepository.RetrieveByEmail(user.Email, &retrievedUser)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = retrievedUser.ComparePassword(user.Password)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	tr, err := h.TokenRepository.GenerateTokens(retrievedUser.ID)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, tr)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	ctp := r.Context().Value(constants.ContextKey).(types.ContextPayload)

	responses.JSON(w, http.StatusOK, ctp.User)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctp := r.Context().Value(constants.ContextKey).(types.ContextPayload)

	err := h.TokenRepository.DeleteTokenMetadata(ctp.TokenId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mateusrlopez/go-market/constants"
	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/repositories"
	"github.com/mateusrlopez/go-market/types"
	"github.com/mateusrlopez/go-market/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandler struct {
	UserRepository  repositories.UserRepository
	TokenRepository repositories.TokenRepository
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.ValidateRegister()

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	result, err := h.UserRepository.Create(&user)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	tr, err := h.TokenRepository.GenerateTokens(result.InsertedID.(primitive.ObjectID).Hex())

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, tr)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.ValidateLogin()

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	retrievedUser := models.User{}
	err = h.UserRepository.RetrieveByEmail(user.Email, &retrievedUser)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = retrievedUser.ComparePassword(user.Password)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	tr, err := h.TokenRepository.GenerateTokens(retrievedUser.ID.Hex())

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, tr)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	ctp := r.Context().Value(constants.ContextKey).(types.ContextPayload)

	utils.JSONResponse(w, http.StatusOK, ctp.User)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctp := r.Context().Value(constants.ContextKey).(types.ContextPayload)

	err := h.TokenRepository.DeleteTokenMetadata(ctp.TokenId)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, nil)
}

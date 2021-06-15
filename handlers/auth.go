package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/repositories"
	"github.com/mateusrlopez/go-market/responses"
	"github.com/mateusrlopez/go-market/utils"
)

type AuthHandler struct {
	Repository repositories.UserRepository
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

	err = h.Repository.Create(&user)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responses.JSON(w, http.StatusCreated, user)
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
	err = h.Repository.RetrieveByEmail(user.Email, &retrievedUser)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = retrievedUser.ComparePassword(user.Password)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := utils.GenerateToken(retrievedUser.ID)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responses.JSON(w, http.StatusOK, retrievedUser)
}

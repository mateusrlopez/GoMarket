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
	UserRepository  repositories.UserRepository
	AdminRepository repositories.AdminRepository
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

	token, err := utils.GenerateToken(user.ID, false)

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

	token, err := utils.GenerateToken(retrievedUser.ID, false)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responses.JSON(w, http.StatusOK, retrievedUser)
}

func (h *AuthHandler) AdminLogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	admin := models.Admin{}
	err = json.Unmarshal(body, &admin)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = admin.ValidateLogin()

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	retrievedAdmin := models.Admin{}
	err = h.AdminRepository.RetrieveByEmail(admin.Email, &retrievedAdmin)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = retrievedAdmin.ComparePassword(admin.Password)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := utils.GenerateToken(retrievedAdmin.ID, true)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))

	responses.JSON(w, http.StatusOK, retrievedAdmin)
}

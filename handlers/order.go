package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/repositories"
	"github.com/mateusrlopez/go-market/shared/types"
	"github.com/mateusrlopez/go-market/shared/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderHandler struct {
	OrderRepository repositories.OrderRepository
}

func (h OrderHandler) Index(w http.ResponseWriter, r *http.Request) {
	query := types.OrderIndexQuery{}

	if err := utils.DecodeQuery(&query, r.URL.Query()); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	orders, err := h.OrderRepository.RetrieveAll(&query)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, orders)
}

func (h OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	order := models.Order{}

	if err = json.Unmarshal(body, &order); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = order.Validate(); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	result, err := h.OrderRepository.Create(&order)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{"id": result.InsertedID.(primitive.ObjectID).Hex()})
}

func (h OrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	order := models.Order{}

	if err = h.OrderRepository.RetrieveByID(id, &order); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, order)
}

func (h OrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	order := models.Order{}

	if err = json.Unmarshal(body, &order); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = order.Validate(); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = h.OrderRepository.Update(id, &order)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, map[string]interface{}{})
}

func (h OrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = h.OrderRepository.Delete(id)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, map[string]interface{}{})
}

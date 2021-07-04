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

type SkuHandler struct {
	SkuRepository repositories.SkuRepository
}

func (h *SkuHandler) Index(w http.ResponseWriter, r *http.Request) {
	query := types.SkuIndexQuery{}

	if err := utils.DecodeQuery(&query, r.URL.Query()); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	skus, err := h.SkuRepository.RetrieveAll(&query)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, skus)
}

func (h *SkuHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	sku := models.SKU{}

	if err = json.Unmarshal(body, &sku); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = sku.Validate(); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	result, err := h.SkuRepository.Create(&sku)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{"id": result.InsertedID.(primitive.ObjectID).Hex()})
}

func (h *SkuHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	sku := models.SKU{}

	if err = json.Unmarshal(body, &sku); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = sku.Validate(); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = h.SkuRepository.Update(id, &sku)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, map[string]interface{}{})
}

func (h *SkuHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = h.SkuRepository.Delete(id)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusNoContent, map[string]interface{}{})
}

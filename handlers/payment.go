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

type PaymentHandler struct {
	PaymentRepository repositories.PaymentRepository
}

func (h PaymentHandler) Index(w http.ResponseWriter, r *http.Request) {
	query := types.PaymentIndexQuery{}

	if err := utils.DecodeQuery(&query, r.URL.Query()); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	payments, err := h.PaymentRepository.RetrieveAll(&query)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, payments)
}

func (h PaymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	payment := models.Payment{}

	if err = json.Unmarshal(body, &payment); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = payment.Validate(); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	stripeResult, err := h.PaymentRepository.StripeCreate(&payment)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	payment.Status = string(stripeResult.Status)
	payment.GatewayID = stripeResult.ID

	result, err := h.PaymentRepository.Create(&payment)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{"id": result.InsertedID.(primitive.ObjectID).Hex()})
}

func (h PaymentHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	payment := models.Payment{}

	if err = h.PaymentRepository.RetrieveByID(id, &payment); err != nil {
		utils.ErrorResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.JSONResponse(w, http.StatusOK, payment)
}

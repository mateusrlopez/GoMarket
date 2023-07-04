package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	customerrors "github.com/mateusrlopez/go-market/internal/errors"
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/services"
)

type PaymentIntentsController struct {
	service services.PaymentIntentsService
}

func NewPaymentIntentsController(service services.PaymentIntentsService) PaymentIntentsController {
	return PaymentIntentsController{
		service: service,
	}
}

func (c PaymentIntentsController) Create(ctx *gin.Context) {
	var input inputs.CreatePaymentIntentInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	secret, err := c.service.Create(input)

	if err != nil {
		if errors.Is(err, customerrors.ErrOrderNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"secret": secret})
}

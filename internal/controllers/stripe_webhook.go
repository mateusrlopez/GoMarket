package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusrlopez/go-market/internal/configurations"
	"github.com/mateusrlopez/go-market/internal/services"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/webhook"
)

type StripeWebhookController struct {
	configuration *configurations.StripeConfiguration
	service       services.PaymentIntentsService
}

func NewStripeWebhookController(configuration *configurations.StripeConfiguration, service services.PaymentIntentsService) StripeWebhookController {
	return StripeWebhookController{
		configuration: configuration,
		service:       service,
	}
}

func (c StripeWebhookController) HandleEvents(ctx *gin.Context) {
	header := ctx.Request.Header.Get("Stripe-Signature")

	payload, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event := stripe.Event{}

	if err := json.Unmarshal(payload, &event); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err = webhook.ConstructEvent(payload, header, c.configuration.WebhookSecret)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch event.Type {
	case "payment_intent.processing":
		var intent stripe.PaymentIntent

		if err := json.Unmarshal(event.Data.Raw, &intent); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gatewayId := intent.ID
		paymentMethodType := string(intent.PaymentMethod.Type)

		if err := c.service.SetOneAsProcessingByGatewayID(gatewayId, paymentMethodType); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{})
	case "payment_intent.succeeded":
		var intent stripe.PaymentIntent

		if err := json.Unmarshal(event.Data.Raw, &intent); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gatewayId := intent.ID

		if err := c.service.SetOneAsSuccessfullByGatewayID(gatewayId); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{})
	case "payment_intent.payment_failed":
		var intent stripe.PaymentIntent

		if err := json.Unmarshal(event.Data.Raw, &intent); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gatewayId := intent.ID

		if err := c.service.SetOneAsFailedByGatewayID(gatewayId); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

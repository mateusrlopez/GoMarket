package services

import (
	"github.com/mateusrlopez/go-market/internal/inputs"
	"github.com/mateusrlopez/go-market/internal/models"
	"github.com/mateusrlopez/go-market/internal/repositories"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
)

type PaymentIntentsService interface {
	Create(input inputs.CreatePaymentIntentInput) (string, error)
	SetOneAsProcessingByGatewayID(gatewayId, paymentMethodType string) error
	SetOneAsSuccessfullByGatewayID(gatewayId string) error
	SetOneAsFailedByGatewayID(gatewayId string) error
}

type paymentIntentsServiceImplementation struct {
	repository    repositories.PaymentIntentsRepository
	ordersService OrdersService
	stripe        *client.API
}

func NewPaymentIntentsService(repository repositories.PaymentIntentsRepository, ordersService OrdersService, stripe *client.API) PaymentIntentsService {
	return paymentIntentsServiceImplementation{
		repository:    repository,
		ordersService: ordersService,
		stripe:        stripe,
	}
}

func (s paymentIntentsServiceImplementation) Create(input inputs.CreatePaymentIntentInput) (string, error) {
	order, err := s.ordersService.FindOneByID(input.OrderID)

	if err != nil {
		return "", err
	}

	intent, err := s.stripe.PaymentIntents.New(&stripe.PaymentIntentParams{
		Amount:   stripe.Int64(order.TotalPrice.Amount),
		Currency: stripe.String(order.TotalPrice.Currency),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	})

	if err != nil {
		return "", err
	}

	payment := models.PaymentIntent{
		OrderID:   input.OrderID,
		Gateway:   "stripe",
		GatewayID: intent.ID,
		Status:    string(intent.Status),
	}

	if err := s.repository.Create(payment); err != nil {
		return "", err
	}

	return intent.ClientSecret, nil
}

func (s paymentIntentsServiceImplementation) SetOneAsProcessingByGatewayID(gatewayId, paymentMethodType string) error {
	return s.repository.SetOneAsProcessingByGatewayID(gatewayId, paymentMethodType)
}

func (s paymentIntentsServiceImplementation) SetOneAsSuccessfullByGatewayID(gatewayId string) error {
	return s.repository.SetOneAsSuccessfullByGatewayID(gatewayId)
}

func (s paymentIntentsServiceImplementation) SetOneAsFailedByGatewayID(gatewayId string) error {
	return s.repository.SetOneAsFailedByGatewayID(gatewayId)
}

package clients

import (
	"github.com/mateusrlopez/go-market/internal/configurations"
	"github.com/stripe/stripe-go/v74/client"
)

func NewStripeClient(configuration *configurations.StripeConfiguration) *client.API {
	return client.New(configuration.Token, nil)
}

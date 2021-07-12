package clients

import (
	"github.com/mateusrlopez/go-market/settings"
	"github.com/stripe/stripe-go/v72/client"
)

func GetStripeClient() *client.API {
	client := &client.API{}
	client.Init(settings.Settings.Stripe.Token, nil)

	return client
}

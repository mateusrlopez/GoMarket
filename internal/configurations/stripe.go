package configurations

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mateusrlopez/go-market/internal/errors"
	"github.com/rs/zerolog/log"
)

type StripeConfiguration struct {
	Token         string `required:"true"`
	WebhookSecret string `required:"true" split_words:"true"`
}

func NewStripeConfiguration() *StripeConfiguration {
	var configuration StripeConfiguration

	if err := envconfig.Process("stripe", &configuration); err != nil {
		log.Fatal().Err(err).Msg(errors.ErrProcessingStripeConfiguration.Error())
	}

	return &configuration
}

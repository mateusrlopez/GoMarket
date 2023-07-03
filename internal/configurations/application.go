package configurations

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mateusrlopez/go-market/internal/errors"
	"github.com/rs/zerolog/log"
)

type ApplicationConfiguration struct {
	Environment string `default:"development"`
}

func NewApplicationConfiguration() *ApplicationConfiguration {
	var configuration ApplicationConfiguration

	if err := envconfig.Process("", &configuration); err != nil {
		log.Fatal().Err(err).Msg(errors.ErrProcessingApplicationConfiguration.Error())
	}

	return &configuration
}

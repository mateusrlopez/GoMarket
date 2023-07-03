package configurations

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mateusrlopez/go-market/internal/errors"
	"github.com/rs/zerolog/log"
)

type ServerConfiguration struct {
	Port int `required:"true"`
}

func NewServerConfiguration() *ServerConfiguration {
	var configuration ServerConfiguration

	if err := envconfig.Process("server", &configuration); err != nil {
		log.Fatal().Err(err).Msg(errors.ErrProcessingServerConfiguration.Error())
	}

	return &configuration
}

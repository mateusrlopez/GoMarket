package configurations

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mateusrlopez/go-market/internal/errors"
	"github.com/rs/zerolog/log"
)

type JwtConfiguration struct {
	Secret string `required:"true"`
}

func NewJwtConfiguration() *JwtConfiguration {
	var configuration JwtConfiguration

	if err := envconfig.Process("jwt", &configuration); err != nil {
		log.Fatal().Err(err).Msg(errors.ErrProcessingJWTConfiguration.Error())
	}

	return &configuration
}

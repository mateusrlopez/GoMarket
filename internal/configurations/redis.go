package configurations

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mateusrlopez/go-market/internal/errors"
	"github.com/rs/zerolog/log"
)

type RedisConfiguration struct {
	Address  string `required:"true"`
	Password string `required:"true"`
	Database int    `required:"true"`
}

func NewRedisConfiguration() *RedisConfiguration {
	var configuration RedisConfiguration

	if err := envconfig.Process("redis", &configuration); err != nil {
		log.Fatal().Err(err).Msg(errors.ErrProcessingRedisConfiguration.Error())
	}

	return &configuration
}

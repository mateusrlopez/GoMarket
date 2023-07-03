package configurations

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mateusrlopez/go-market/internal/errors"
	"github.com/rs/zerolog/log"
)

type MongoConfiguration struct {
	Uri      string `required:"true"`
	Database string `required:"true"`
}

func NewMongoConfiguration() *MongoConfiguration {
	var configuration MongoConfiguration

	if err := envconfig.Process("mongo", &configuration); err != nil {
		log.Fatal().Err(err).Msg(errors.ErrProcessingMongoConfiguration.Error())
	}

	return &configuration
}

package clients

import (
	"context"

	"github.com/mateusrlopez/go-market/internal/configurations"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func NewRedisClient(configuration *configurations.RedisConfiguration) *redis.Client {
	client := redis.NewClient(&redis.Options{Addr: configuration.Address, Password: configuration.Password, DB: configuration.Database})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal().Err(err).Msg("could not ping the Redis instance")
	}

	return client
}

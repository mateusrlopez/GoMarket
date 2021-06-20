package database

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/mateusrlopez/go-market/settings"
	log "github.com/sirupsen/logrus"
)

func GetRedisConnection() *redis.Client {
	opt, err := redis.ParseURL(formatRedisURL())

	if err != nil {
		log.Fatalf("Error opening connection with redis database: %s", err)
		return nil
	}

	return redis.NewClient(opt)
}

func formatRedisURL() string {
	return fmt.Sprintf("redis://%s:%d/%d",
		settings.Settings.Redis.Host,
		settings.Settings.Redis.Port,
		settings.Settings.Redis.Database,
	)
}

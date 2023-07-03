package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenRepository interface {
	SaveTokenMetadata(userId, tokenId string, expiration time.Time) error
	RetrieveTokenMetadata(userId string) (string, error)
	DeleteTokenMetadata(userId string) error
}

type redisTokenRepository struct {
	redis *redis.Client
}

func NewTokenRepository(redis *redis.Client) TokenRepository {
	return redisTokenRepository{
		redis: redis,
	}
}

func (r redisTokenRepository) SaveTokenMetadata(userId, tokenId string, expiration time.Time) error {
	return r.redis.Set(context.Background(), fmt.Sprintf("token:%s", userId), tokenId, time.Until(expiration)).Err()
}

func (r redisTokenRepository) RetrieveTokenMetadata(userId string) (string, error) {
	result, err := r.redis.Get(context.Background(), fmt.Sprintf("token:%s", userId)).Result()

	if err != nil {
		return "", err
	}

	return result, nil
}

func (r redisTokenRepository) DeleteTokenMetadata(userId string) error {
	return r.redis.Del(context.Background(), fmt.Sprintf("token:%s", userId)).Err()
}

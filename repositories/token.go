package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"github.com/mateusrlopez/go-market/settings"
	"github.com/mateusrlopez/go-market/types"
	"github.com/twinj/uuid"
)

type TokenRepository struct {
	DB *redis.Client
}

func (r *TokenRepository) GenerateTokens(sub string) (*types.TokensReturn, error) {
	accessClaims := jwt.MapClaims{}

	accessUuid := uuid.NewV4().String()
	accessExp := time.Now().Add(time.Hour * 1).Unix()

	accessClaims["jti"] = accessUuid
	accessClaims["sub"] = sub
	accessClaims["exp"] = accessExp

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(settings.Settings.Server.AccessSecret))

	if err != nil {
		return nil, err
	}

	err = r.StoreTokenMetadata(accessUuid, accessExp, sub)

	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{}

	refreshUuid := uuid.NewV4().String()
	refreshExp := time.Now().Add(time.Hour * 24).Unix()

	refreshClaims["jti"] = refreshUuid
	refreshClaims["sub"] = sub
	refreshClaims["exp"] = refreshExp

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(settings.Settings.Server.RefreshSecret))

	if err != nil {
		return nil, err
	}

	err = r.StoreTokenMetadata(refreshUuid, refreshExp, sub)

	if err != nil {
		return nil, err
	}

	return &types.TokensReturn{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (r *TokenRepository) StoreTokenMetadata(uuid string, exp int64, sub string) error {
	return r.DB.Set(context.Background(), uuid, sub, time.Until(time.Unix(exp, 0))).Err()
}

func (r *TokenRepository) ValidateToken(tokenString string, secret string) (*types.TokenMetadataReturn, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return &types.TokenMetadataReturn{UUID: claims["jti"].(string), UserId: claims["sub"].(string)}, nil
}

func (r *TokenRepository) RetrieveTokenMetadata(uuid string) error {
	return r.DB.Get(context.Background(), uuid).Err()
}

func (r *TokenRepository) DeleteTokenMetadata(uuid string) error {
	return r.DB.Del(context.Background(), uuid).Err()
}

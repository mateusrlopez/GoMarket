package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mateusrlopez/go-market/internal/configurations"
	"github.com/mateusrlopez/go-market/internal/repositories"
)

type TokenService interface {
	AssignToken(userId string) (string, error)
	ValidateToken(str string) (string, error)
	UnassignToken(userId string) error
}

type tokenServiceImplementation struct {
	configuration *configurations.JwtConfiguration
	repository    repositories.TokenRepository
}

func NewTokenService(configuration *configurations.JwtConfiguration, repository repositories.TokenRepository) TokenService {
	return tokenServiceImplementation{
		configuration: configuration,
		repository:    repository,
	}
}

func (s tokenServiceImplementation) AssignToken(userId string) (string, error) {
	key := s.configuration.Secret

	tokenId := uuid.NewString()
	expiration := time.Now().Add(30 * time.Minute)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": userId, "jti": tokenId, "exp": expiration.Unix()}).SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	if err = s.repository.SaveTokenMetadata(userId, tokenId, expiration); err != nil {
		return "", err
	}

	return token, nil
}

func (s tokenServiceImplementation) ValidateToken(str string) (string, error) {
	token, err := jwt.Parse(str, func(t *jwt.Token) (interface{}, error) {
		if method := t.Method.Alg(); method != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("unexpected signing method: %s", method)
		}

		return []byte(s.configuration.Secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	userId := claims["sub"].(string)
	tokenId := claims["jti"].(string)

	retrievedId, err := s.repository.RetrieveTokenMetadata(userId)

	if err != nil {
		return "", err
	}

	if tokenId != retrievedId {
		return "", fmt.Errorf("received token's id does not match the one associated with the received token's user")
	}

	return userId, nil
}

func (s tokenServiceImplementation) UnassignToken(userId string) error {
	return s.repository.DeleteTokenMetadata(userId)
}

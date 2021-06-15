package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mateusrlopez/go-market/settings"
)

func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}

	claims["uid"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(settings.Settings.Server.Secret))
}

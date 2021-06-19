package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mateusrlopez/go-market/settings"
)

func GenerateToken(sub uint, adm bool) (string, error) {
	claims := jwt.MapClaims{}

	claims["adm"] = adm
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(settings.Settings.Server.Secret))
}

func ValidateToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(settings.Settings.Server.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return map[string]interface{}{"admin": claims["adm"], "userId": claims["sub"]}, nil
}

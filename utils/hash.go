package utils

import "golang.org/x/crypto/bcrypt"

func Hash(value string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
}

func CompareHash(hashedValue string, plainValue string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedValue), []byte(plainValue))
}

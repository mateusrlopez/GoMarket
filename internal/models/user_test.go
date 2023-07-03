package models

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestComparePassword(t *testing.T) {
	t.Run("NominalCase", func(t *testing.T) {
		password := faker.Password()

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		user := User{Password: string(hash)}

		err = user.ComparePassword(password)

		assert.NoError(t, err)
	})

	t.Run("ErrorCase", func(t *testing.T) {
		password := faker.Password()

		user := User{Password: password}

		err := user.ComparePassword(password)

		assert.Error(t, err)
	})
}

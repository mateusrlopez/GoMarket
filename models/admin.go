package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mateusrlopez/go-market/utils"
)

type Admin struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;" json:"id"`
	FirstName string    `gorm:"size:255;not null;" json:"first_name"`
	LastName  string    `gorm:"size:255;not null;" json:"last_name"`
	Email     string    `gorm:"size:255;not null;unique;" json:"email"`
	Password  string    `gorm:"size:255;not null;" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Admin) ComparePassword(password string) error {
	return utils.CompareHash(a.Password, password)
}

func (a *Admin) ValidateLogin() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Email, validation.Required),
		validation.Field(&a.Password, validation.Required),
	)
}

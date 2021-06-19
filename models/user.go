package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mateusrlopez/go-market/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;" json:"id"`
	FirstName string    `gorm:"size:255;not null;" json:"first_name"`
	LastName  string    `gorm:"size:255;not null;" json:"last_name"`
	Email     string    `gorm:"size:255;not null;unique;" json:"email"`
	Password  string    `gorm:"size:255;not null;" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := utils.Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return utils.CompareHash(u.Password, password)
}

func (u *User) ValidateRegister() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required),
	)
}

func (u *User) ValidateLogin() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Password, validation.Required),
	)
}

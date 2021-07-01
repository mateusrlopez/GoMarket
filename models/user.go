package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mateusrlopez/go-market/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password"`
	Birthdate string             `json:"birthdate" bson:"birthdate"`
	Admin     bool               `json:"admin" bson:"admin"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
}

func (u *User) BeforeInsert() error {
	u.CreatedAt = time.Now()
	hashedPassword, err := utils.Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u User) ComparePassword(password string) error {
	return utils.CompareHash(u.Password, password)
}

func (u User) ValidateRegister() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.FirstName, validation.Required),
		validation.Field(&u.LastName, validation.Required),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Birthdate, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&u.Password, validation.Required),
	)
}

func (u User) ValidateLogin() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Password, validation.Required),
	)
}

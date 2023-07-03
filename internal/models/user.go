package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"-" bson:"password,omitempty"`
	FirstName string    `json:"firstName" bson:"firstName"`
	LastName  string    `json:"lastName" bson:"lastName"`
	Addresses []Address `json:"addresses" bson:"addresses"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
}

func (u User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

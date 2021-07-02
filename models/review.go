package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ProductID  primitive.ObjectID `json:"productId" bson:"productId"`
	Name       string             `json:"name" bson:"name"`
	Rating     float64            `json:"rating" bson:"rating"`
	Commentary string             `json:"commentary" bson:"commentary"`
	CreatedAt  time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
}

func (r *Review) BeforeInsert() {
	r.CreatedAt = time.Now()
}

func (r Review) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.ProductID, validation.Required),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Rating, validation.Required, validation.Max(5.0)),
		validation.Field(&r.Commentary, validation.Required),
	)
}

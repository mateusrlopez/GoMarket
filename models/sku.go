package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mateusrlopez/go-market/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SKU struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ProductID       primitive.ObjectID `json:"productId" bson:"productId"`
	SKU             string             `json:"sku" bson:"sku"`
	Price           types.Price        `json:"price" bson:"price"`
	Available       bool               `json:"available" bson:"available"`
	QuantityInStock int                `json:"quantityInStock" bson:"quantityInStock"`
	Images          []string           `json:"images" bson:"images"`
	Attributes      []types.Attributes `json:"attributes" bson:"attributes"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
}

func (s *SKU) BeforeInsert() {
	s.CreatedAt = time.Now()
}

func (s SKU) Validate() error {
	return validation.ValidateStruct(
		&s,
		validation.Field(&s.ProductID, validation.Required, is.MongoID),
		validation.Field(&s.SKU, validation.Required),
		validation.Field(&s.Price, validation.Required),
		validation.Field(&s.Available, validation.Required),
		validation.Field(&s.QuantityInStock, validation.Required),
		validation.Field(&s.Images, validation.Required, validation.Each(is.URL)),
		validation.Field(&s.Attributes, validation.NotNil),
	)
}

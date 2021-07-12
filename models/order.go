package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mateusrlopez/go-market/shared/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"userId" bson:"userId"`
	Shipping      types.Shipping     `json:"shipping" bson:"shipping"`
	OrderItems    []types.OrderItem  `json:"orderItems" bson:"orderItems"`
	ItemsPrice    types.Price        `json:"itemsPrice" bson:"itemsPrice"`
	ShippingPrice types.Price        `json:"shippingPrice" bson:"shippingPrice"`
	Tax           types.Price        `json:"tax" bson:"tax"`
	TotalPrice    types.Price        `json:"totalPrice" bson:"totalPrice"`
	CreatedAt     time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
}

func (o *Order) BeforeInsert() {
	o.CreatedAt = time.Now()
}

func (o Order) Validate() error {
	return validation.ValidateStruct(
		&o,
		validation.Field(&o.UserID, validation.Required),
		validation.Field(&o.Shipping, validation.Required),
		validation.Field(&o.OrderItems, validation.Required),
		validation.Field(&o.ItemsPrice, validation.Required),
		validation.Field(&o.ShippingPrice, validation.Required),
		validation.Field(&o.Tax, validation.Required),
		validation.Field(&o.TotalPrice, validation.Required),
	)
}

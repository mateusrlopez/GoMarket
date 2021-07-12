package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attributes struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type Options struct {
	Name   string   `json:"name" bson:"name"`
	Values []string `json:"values" bson:"values"`
}

type Price struct {
	Value    float64 `json:"value" bson:"value"`
	Currency string  `json:"currency" bson:"currency"`
}

type Shipping struct {
	Address string `json:"address" bson:"address"`
	City    string `json:"city" bson:"city"`
	ZipCode string `json:"zipcode" bson:"zipcode"`
	Country string `json:"country" bson:"country"`
}

type OrderItem struct {
	Name       string             `json:"name" bson:"name"`
	Image      string             `json:"image" bson:"image"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	Price      Price              `json:"price" bson:"price"`
	Attributes []Attributes       `json:"attributes" bson:"attributes"`
	SkuID      primitive.ObjectID `json:"skuId" bson:"skuId"`
}

func (a Attributes) Validate() error {
	return validation.ValidateStruct(
		&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Value, validation.Required),
	)
}

func (o Options) Validate() error {
	return validation.ValidateStruct(
		&o,
		validation.Field(&o.Name, validation.Required),
		validation.Field(&o.Values, validation.Required),
	)
}

func (p Price) Validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.Value, validation.Required),
		validation.Field(&p.Currency, validation.Required),
	)
}

func (s Shipping) Validate() error {
	return validation.ValidateStruct(
		&s,
		validation.Field(&s.Address, validation.Required),
		validation.Field(&s.City, validation.Required),
		validation.Field(&s.ZipCode, validation.Required),
		validation.Field(&s.Country, validation.Required),
	)
}

func (oi OrderItem) Validate() error {
	return validation.ValidateStruct(
		&oi,
		validation.Field(&oi.Name, validation.Required),
		validation.Field(&oi.Image, validation.Required, is.URL),
		validation.Field(&oi.Quantity, validation.Required),
		validation.Field(&oi.Price, validation.Required),
		validation.Field(&oi.Attributes, validation.NotNil),
		validation.Field(&oi.SkuID, validation.Required),
	)
}

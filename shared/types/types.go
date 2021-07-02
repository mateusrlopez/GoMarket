package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContextKey string

type TokensReturn struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenMetadataReturn struct {
	UUID   string
	UserId string
}

type Attributes struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type ReviewIndexQuery struct {
	ProductID primitive.ObjectID `schema:"productId" bson:"productId"`
}

type Options struct {
	Name   string   `json:"name" bson:"name"`
	Values []string `json:"values" bson:"values"`
}

type Price struct {
	Value    float64 `json:"value" bson:"value"`
	Currency string  `json:"currency" bson:"currency"`
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

package types

import validation "github.com/go-ozzo/ozzo-validation/v4"

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

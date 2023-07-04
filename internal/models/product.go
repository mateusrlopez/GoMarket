package models

import (
	"time"
)

type ProductSpecification struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type ProductCustomization struct {
	Name    string                       `json:"name" bson:"name"`
	Options []ProductCustomizationOption `json:"options" bson:"options"`
}

type ProductCustomizationOption struct {
	Name          string `json:"name" bson:"name"`
	PriceIncrease Money  `json:"priceIncrease" bson:"priceIncrease"`
}

type Product struct {
	ID             string                 `json:"id" bson:"_id,omitempty"`
	Name           string                 `json:"name" bson:"name"`
	Type           string                 `json:"type" bson:"type"`
	Specifications []ProductSpecification `json:"specifications" bson:"specifications"`
	Customizations []ProductCustomization `json:"customizations" bson:"customizations"`
	BasePrice      Money                  `json:"basePrice" bson:"basePrice"`
	CreatedAt      time.Time              `json:"createdAt" bson:"createdAt,omitempty"`
}

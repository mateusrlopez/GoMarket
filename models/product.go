package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SKU struct {
	SKU             string  `json:"sku" bson:"sku"`
	Size            string  `json:"size" bson:"size"`
	Color           string  `json:"color" bson:"color"`
	Price           float64 `json:"price" bson:"price"`
	Image           string  `json:"image" bson:"image"`
	QuantityInStock int     `json:"quantityInStock" bson:"quantityInStock"`
}

type Product struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name"`
	Type           string             `json:"type" bson:"type"`
	Description    string             `json:"description" bson:"description"`
	BasePrice      float64            `json:"basePrice" bson:"basePrice"`
	Brand          string             `json:"brand" bson:"brand"`
	Gender         string             `json:"gender" bson:"gender"`
	ReleaseDate    string             `json:"releaseDate" bson:"releaseDate"`
	Release        bool               `json:"release" bson:"release"`
	Available      bool               `json:"available" bson:"available"`
	Material       string             `json:"material" bson:"material"`
	ThumbnailImage string             `json:"thumbnailImage" bson:"thumbnailImage"`
	Images         []string           `json:"images" bson:"images"`
	Categories     []string           `json:"categories" bson:"categories"`
	Sizes          []string           `json:"sizes" bson:"sizes"`
	Colors         []string           `json:"colors" bson:"colors"`
	SKUs           []SKU              `json:"skus" bson:"skus"`
	CreatedAt      time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

func (p *Product) BeforeInsert() {
	p.CreatedAt = time.Now()
}

func (p *Product) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Type, validation.Required),
		validation.Field(&p.Description, validation.NotNil),
		validation.Field(&p.BasePrice, validation.Required),
		validation.Field(&p.Brand, validation.Required),
		validation.Field(&p.Gender, validation.Required),
		validation.Field(&p.ReleaseDate, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&p.Release, validation.Required),
		validation.Field(&p.Available, validation.Required),
		validation.Field(&p.Material, validation.Required),
		validation.Field(&p.ThumbnailImage, validation.NotNil, is.URL),
		validation.Field(&p.Images, validation.NotNil, validation.Each(is.URL)),
		validation.Field(&p.Categories, validation.Required),
		validation.Field(&p.Sizes, validation.Required),
		validation.Field(&p.Colors, validation.Required),
		validation.Field(&p.SKUs, validation.NotNil),
	)
}

package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mateusrlopez/go-market/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name"`
	Type           string             `json:"type" bson:"type"`
	Description    string             `json:"description" bson:"description"`
	BasePrice      types.Price        `json:"basePrice" bson:"basePrice"`
	Brand          string             `json:"brand" bson:"brand"`
	ReleaseDate    string             `json:"releaseDate" bson:"releaseDate"`
	ThumbnailImage string             `json:"thumbnailImage" bson:"thumbnailImage"`
	Categories     []string           `json:"categories" bson:"categories"`
	Features       []string           `json:"features" bson:"features"`
	Options        []types.Options    `json:"options" bson:"options"`
	Details        []types.Attributes `json:"details" bson:"details"`
	CreatedAt      time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
}

func (p *Product) BeforeInsert() {
	p.CreatedAt = time.Now()
}

func (p Product) Validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Type, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.BasePrice, validation.Required),
		validation.Field(&p.Brand, validation.Required),
		validation.Field(&p.ReleaseDate, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&p.ThumbnailImage, validation.Required, is.URL),
		validation.Field(&p.Categories, validation.Required),
		validation.Field(&p.Features, validation.NotNil),
		validation.Field(&p.Options, validation.NotNil),
		validation.Field(&p.Details, validation.NotNil),
	)
}

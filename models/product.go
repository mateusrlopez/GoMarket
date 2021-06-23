package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID             primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name           string             `json:"name" bson:"name"`
	Type           string             `json:"type" bson:"type"`
	Description    string             `json:"description" bson:"description"`
	BasePrice      float64            `json:"basePrice" bson:"basePrice"`
	Brand          string             `json:"brand" bson:"brand"`
	Release        bool               `json:"release" bson:"release"`
	Available      bool               `json:"available" bson:"available"`
	Material       string             `json:"material" bson:"material"`
	ThumbnailImage string             `json:"thumbnailImage" bson:"thumbnailImage"`
	Images         []string           `json:"images" bson:"images"`
	Categories     []string           `json:"categories" bson:"categories"`
	Sizes          []string           `json:"sizes" bson:"sizes"`
	Colors         []string           `json:"colors" bson:"colors"`
	CreatedAt      time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

func (p *Product) BeforeInsert() {
	p.CreatedAt = time.Now()
}

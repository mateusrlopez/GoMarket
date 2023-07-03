package models

import "time"

type Review struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Title     string    `json:"title" bson:"title"`
	Comment   string    `json:"comment" bson:"comment"`
	Rating    float64   `json:"rating" bson:"rating"`
	UserID    string    `json:"userId" bson:"userId,omitempty"`
	ProductID string    `json:"productId" bson:"productId,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
}

package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mateusrlopez/go-market/shared/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Price         types.Price        `json:"price" bson:"price"`
	PaymentMethod string             `json:"paymentMethod" bson:"paymentMethod"`
	Source        string             `json:"source" bson:"-"`
	Status        string             `json:"status,omitempty" bson:"status"`
	Gateway       string             `json:"gateway" bson:"gateway"`
	GatewayID     string             `json:"gatewayId,omitempty" bson:"gatewayId"`
	OrderID       primitive.ObjectID `json:"orderId" bson:"orderId"`
	CreatedAt     time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
}

func (p *Payment) BeforeInsert() {
	p.CreatedAt = time.Now()
}

func (p Payment) Validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.Price, validation.Required),
		validation.Field(&p.PaymentMethod, validation.Required),
		validation.Field(&p.Gateway, validation.Required),
		validation.Field(&p.OrderID, validation.Required),
	)
}

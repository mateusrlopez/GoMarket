package models

import "time"

type PaymentIntent struct {
	ID                string    `json:"id" bson:"_id"`
	OrderID           string    `json:"orderId" bson:"orderId"`
	Gateway           string    `json:"-" bson:"gateway"`
	GatewayID         string    `json:"-" bson:"gatewayId"`
	PaymentMethodType string    `json:"paymentMethodType" bson:"paymentMethodType,omitempty"`
	Status            string    `json:"status" bson:"status"`
	CreatedAt         time.Time `json:"createdAt" bson:"createdAt"`
	SucceededAt       time.Time `json:"succeededAt" bson:"succeededAt,omitempty"`
	FailedAt          time.Time `json:"failedAt" bson:"failtedAt,omitempty"`
}

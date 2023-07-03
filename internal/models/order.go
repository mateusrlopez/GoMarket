package models

import (
	"time"

	"github.com/Rhymond/go-money"
)

type OrderItem struct {
	ProductID      string       `json:"productId" bson:"productId"`
	ProductName    string       `json:"productName" bson:"productName"`
	Quantity       uint         `json:"quantity" bson:"quantity"`
	Customizations []string     `json:"customizations" bson:"customizations"`
	Price          *money.Money `json:"price" bson:"price"`
}

type OrderShipping struct {
	RecipientName  string  `json:"recipientName" bson:"recipientName"`
	RecipientPhone string  `json:"recipientPhone" bson:"recipientPhone"`
	Address        Address `json:"address" bson:"address"`
}

type Order struct {
	ID           string        `json:"id" bson:"_id,omitempty"`
	UserID       string        `json:"userId" bson:"userId,omitempty"`
	CurrencyCode string        `json:"currencyCode" bson:"currencyCode"`
	Status       string        `json:"status" bson:"status,omitempty"`
	Shipping     OrderShipping `json:"shipping" bson:"shipping"`
	TotalPrice   *money.Money  `json:"totalPrice" bson:"totalPrice"`
	Items        []OrderItem   `json:"items" bson:"items"`
	CreatedAt    time.Time     `json:"createdAt" bson:"createdAt,omitempty"`
}

func (o *Order) SetTotalPrice() error {
	amount := money.New(0, o.CurrencyCode)

	for _, item := range o.Items {
		var err error

		amount, err = amount.Add(item.Price.Multiply(int64(item.Quantity)))

		if err != nil {
			return err
		}
	}

	o.TotalPrice = amount

	return nil
}

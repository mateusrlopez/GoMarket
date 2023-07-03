package inputs

import "github.com/Rhymond/go-money"

type UpdateOrderInput struct {
	CurrencyCode string `json:"currencyCode" binding:"required"`
	Shipping     struct {
		RecipientName  string `json:"recipientName" binding:"required"`
		RecipientPhone string `json:"recipientPhone" binding:"required"`
		Address        struct {
			City       string `json:"city" binding:"required"`
			State      string `json:"state" binding:"required"`
			Country    string `json:"country" binding:"required"`
			PostalCode string `json:"postalCode" binding:"required"`
			Line1      string `json:"line1" binding:"required"`
			Line2      string `json:"line2"`
		} `json:"addresses" binding:"required"`
	} `json:"shipping" binding:"required"`
	Items []struct {
		ProductID      string       `json:"productId" binding:"required"`
		ProductName    string       `json:"productName" binding:"required"`
		Quantity       uint         `json:"quantity" binding:"required"`
		Customizations []string     `json:"customizations" binding:"required"`
		Price          *money.Money `json:"price" binding:"required"`
	} `json:"items" binding:"required"`
}

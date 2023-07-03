package inputs

import "github.com/Rhymond/go-money"

type UpdateProductInput struct {
	Name           string `json:"name" binding:"required"`
	Type           string `json:"type" binding:"required"`
	Specifications []struct {
		Name  string `json:"name" binding:"required"`
		Value string `json:"value" binding:"required"`
	} `json:"specifications"`
	Customizations []struct {
		Name    string `json:"name" binding:"required"`
		Options []struct {
			Name          string       `json:"name" binding:"required"`
			PriceIncrease *money.Money `json:"priceIncrease" binding:"required"`
		} `json:"options" binding:"required"`
	} `json:"customizations"`
	Price *money.Money `json:"price" binding:"required"`
}

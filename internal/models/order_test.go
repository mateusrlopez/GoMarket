package models

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestSetTotalPrice(t *testing.T) {
	t.Run("NominalCase", func(t *testing.T) {
		order := Order{
			CurrencyCode: money.USD,
			Items: []OrderItem{
				{
					Price:    money.New(100, money.USD),
					Quantity: 3,
				},
			},
		}

		err := order.SetTotalPrice()

		assert.NoError(t, err)
		assert.Equal(t, int64(300), order.TotalPrice.Amount())
		assert.Equal(t, money.USD, order.TotalPrice.Currency().Code)
	})

	t.Run("ErrorCase", func(t *testing.T) {
		order := Order{
			CurrencyCode: money.USD,
			Items: []OrderItem{
				{
					Price:    money.New(100, money.EUR),
					Quantity: 3,
				},
			},
		}

		err := order.SetTotalPrice()

		assert.Error(t, err)
		assert.Nil(t, order.TotalPrice)
	})
}

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetTotalPrice(t *testing.T) {
	t.Run("NominalCase", func(t *testing.T) {
		order := Order{
			CurrencyCode: "USD",
			Items: []OrderItem{
				{
					Price:    NewMoney(100, "USD"),
					Quantity: 3,
				},
			},
		}

		err := order.SetTotalPrice()

		assert.NoError(t, err)
		assert.Equal(t, int64(300), order.TotalPrice.Amount)
		assert.Equal(t, "USD", order.TotalPrice.Currency)
	})

	t.Run("ErrorCase", func(t *testing.T) {
		order := Order{
			CurrencyCode: "USD",
			Items: []OrderItem{
				{
					Price:    NewMoney(100, "EUR"),
					Quantity: 3,
				},
			},
		}

		err := order.SetTotalPrice()

		assert.Error(t, err)
	})
}

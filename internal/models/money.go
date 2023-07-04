package models

import "errors"

type Money struct {
	Amount   int64  `json:"amount" bson:"amount"`
	Currency string `json:"currency" bson:"currency"`
}

func NewMoney(amount int64, currency string) Money {
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

func (m Money) Add(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, errors.New("currency codes do not match")
	}

	m.Amount += other.Amount

	return m, nil
}

func (m Money) Multiply(value int64) Money {
	m.Amount *= value

	return m
}

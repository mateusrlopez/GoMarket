package types

import "github.com/mateusrlopez/go-market/models"

type ContextKey string

type ContextPayload struct {
	User    *models.User
	TokenId string
}

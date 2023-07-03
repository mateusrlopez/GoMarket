package errors

import "errors"

var (
	ErrUserNotFound    = errors.New("user with given parameters was not found")
	ErrProductNotFound = errors.New("product with given parameters was not found")
	ErrReviewNotFound  = errors.New("review with given parameters was not found")
	ErrOrderNotFound   = errors.New("order with given parameters was not found")
)

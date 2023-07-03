package errors

import "errors"

var (
	ErrProcessingApplicationConfiguration = errors.New("could not process the application configuration")
	ErrProcessingJWTConfiguration         = errors.New("could not process the JWT configuration")
	ErrProcessingMongoConfiguration       = errors.New("could not process the MongoDB configuration")
	ErrProcessingRedisConfiguration       = errors.New("could not process the Redis configuration")
	ErrProcessingServerConfiguration      = errors.New("could not process the server configuration")
	ErrProcessingStripeConfiguration      = errors.New("could not process the Stripe configuration")
)

package clients

import (
	"context"

	"github.com/mateusrlopez/go-market/internal/configurations"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(configuration *configurations.MongoConfiguration) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(configuration.Uri))

	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to the MongoDB instance")
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal().Err(err).Msg("could not ping the MongoDB instance")
	}

	return client.Database(configuration.Database)
}

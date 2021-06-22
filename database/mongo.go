package database

import (
	"context"
	"fmt"

	"github.com/mateusrlopez/go-market/settings"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoConnection() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(formatMongoURI()))

	if err != nil {
		log.Fatalf("Error opening connection with mongo database: %s", err)
		return nil
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatalf("Error pinging the mongo database: %s", err)
		return nil
	}

	return client.Database(settings.Settings.Database.Name)
}

func formatMongoURI() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d",
		settings.Settings.Database.UserName,
		settings.Settings.Database.Password,
		settings.Settings.Database.Host,
		settings.Settings.Database.Port,
	)
}

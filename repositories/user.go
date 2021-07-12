package repositories

import (
	"context"

	"github.com/mateusrlopez/go-market/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	Collection *mongo.Collection
}

func (r UserRepository) Create(user *models.User) (*mongo.InsertOneResult, error) {
	err := user.BeforeInsert()

	if err != nil {
		return nil, err
	}

	return r.Collection.InsertOne(context.Background(), user)
}

func (r UserRepository) RetrieveByEmail(email string, user *models.User) error {
	return r.Collection.FindOne(context.Background(), bson.M{"email": bson.M{"$eq": email}}).Decode(user)
}

func (r UserRepository) RetriveByID(id primitive.ObjectID, user *models.User) error {
	return r.Collection.FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, options.FindOne().SetProjection(bson.M{"password": 0})).Decode(user)
}

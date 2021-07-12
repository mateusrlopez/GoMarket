package repositories

import (
	"context"

	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	Collection *mongo.Collection
}

func (r OrderRepository) Create(order *models.Order) (*mongo.InsertOneResult, error) {
	order.BeforeInsert()

	return r.Collection.InsertOne(context.Background(), order)
}

func (r OrderRepository) RetrieveAll(filter *types.OrderIndexQuery) ([]models.Order, error) {
	var orders []models.Order

	cursor, err := r.Collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r OrderRepository) RetrieveByID(id primitive.ObjectID, order *models.Order) error {
	return r.Collection.FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Decode(order)
}

func (r OrderRepository) Update(id primitive.ObjectID, order *models.Order) (*mongo.UpdateResult, error) {
	return r.Collection.UpdateOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": order})
}

func (r OrderRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return r.Collection.DeleteOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}})
}

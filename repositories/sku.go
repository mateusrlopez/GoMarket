package repositories

import (
	"context"

	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SkuRepository struct {
	Collection *mongo.Collection
}

func (r SkuRepository) Create(sku *models.SKU) (*mongo.InsertOneResult, error) {
	sku.BeforeInsert()

	return r.Collection.InsertOne(context.Background(), sku)
}

func (r SkuRepository) RetrieveAll(filter *types.SkuIndexQuery) ([]models.SKU, error) {
	var skus []models.SKU

	cursor, err := r.Collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &skus); err != nil {
		return nil, err
	}

	return skus, nil
}

func (r SkuRepository) Update(id primitive.ObjectID, sku *models.SKU) (*mongo.UpdateResult, error) {
	return r.Collection.UpdateOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": sku})
}

func (r SkuRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return r.Collection.DeleteOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}})
}

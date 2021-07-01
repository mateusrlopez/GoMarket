package repositories

import (
	"context"

	"github.com/mateusrlopez/go-market/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	Collection *mongo.Collection
}

func (r *ProductRepository) Create(product *models.Product) (*mongo.InsertOneResult, error) {
	product.BeforeInsert()

	return r.Collection.InsertOne(context.Background(), product)
}

func (r *ProductRepository) RetrieveAll() ([]models.Product, error) {
	var products []models.Product

	cursor, err := r.Collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	if cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) RetrieveByID(id primitive.ObjectID, product *models.Product) error {
	return r.Collection.FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Decode(product)
}

func (r *ProductRepository) Update(id primitive.ObjectID, product *models.Product) (*mongo.UpdateResult, error) {
	return r.Collection.UpdateOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": product})
}

func (r *ProductRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return r.Collection.DeleteOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}})
}

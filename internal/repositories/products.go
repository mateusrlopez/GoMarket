package repositories

import (
	"context"
	"errors"
	"time"

	customerrors "github.com/mateusrlopez/go-market/internal/errors"
	"github.com/mateusrlopez/go-market/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductsRepository interface {
	Create(product models.Product) (models.Product, error)
	FindMany() ([]models.Product, error)
	FindOneByID(id string) (models.Product, error)
	UpdateOneByID(id string, data models.Product) (models.Product, error)
	DeleteOneByID(id string) error
}

type mongoProductsRepository struct {
	db *mongo.Database
}

func NewProductsRepository(db *mongo.Database) ProductsRepository {
	return mongoProductsRepository{
		db: db,
	}
}

func (r mongoProductsRepository) Create(product models.Product) (models.Product, error) {
	product.ID = primitive.NewObjectID().Hex()
	product.CreatedAt = time.Now()

	_, err := r.db.Collection("products").InsertOne(context.Background(), &product)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r mongoProductsRepository) FindMany() ([]models.Product, error) {
	var products []models.Product

	cursor, err := r.db.Collection("products").Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var product models.Product

		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r mongoProductsRepository) FindOneByID(id string) (models.Product, error) {
	var product models.Product

	if err := r.db.Collection("products").FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&product); err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (r mongoProductsRepository) UpdateOneByID(id string, data models.Product) (models.Product, error) {
	var updated models.Product

	after := options.After

	if err := r.db.Collection("products").FindOneAndUpdate(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": &data}, &options.FindOneAndUpdateOptions{ReturnDocument: &after}).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Product{}, customerrors.ErrProductNotFound
		}

		return models.Product{}, err
	}

	return updated, nil
}

func (r mongoProductsRepository) DeleteOneByID(id string) error {
	if err := r.db.Collection("products").FindOneAndDelete(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return customerrors.ErrProductNotFound
		}

		return err
	}

	return nil
}

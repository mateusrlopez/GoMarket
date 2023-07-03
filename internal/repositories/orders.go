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

type OrdersRepository interface {
	Create(order models.Order) (models.Order, error)
	FindMany(filter models.Order) ([]models.Order, error)
	FindOneByID(id string) (models.Order, error)
	UpdateOneByID(id string, data models.Order) (models.Order, error)
}

type mongoOrdersRepository struct {
	db *mongo.Database
}

func NewOrdersRepository(db *mongo.Database) OrdersRepository {
	return mongoOrdersRepository{
		db: db,
	}
}

func (r mongoOrdersRepository) Create(order models.Order) (models.Order, error) {
	order.ID = primitive.NewObjectID().Hex()
	order.Status = "Expecting Payment"
	order.CreatedAt = time.Now()

	if err := order.SetTotalPrice(); err != nil {
		return models.Order{}, err
	}

	_, err := r.db.Collection("orders").InsertOne(context.Background(), &order)

	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (r mongoOrdersRepository) FindMany(filter models.Order) ([]models.Order, error) {
	var orders []models.Order

	cursor, err := r.db.Collection("orders").Find(context.Background(), &filter)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var order models.Order

		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (r mongoOrdersRepository) FindOneByID(id string) (models.Order, error) {
	var order models.Order

	if err := r.db.Collection("orders").FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&order); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Order{}, customerrors.ErrOrderNotFound
		}

		return models.Order{}, err
	}

	return order, nil
}

func (r mongoOrdersRepository) UpdateOneByID(id string, data models.Order) (models.Order, error) {
	var updated models.Order

	if err := data.SetTotalPrice(); err != nil {
		return models.Order{}, err
	}

	after := options.After

	if err := r.db.Collection("orders").FindOneAndUpdate(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": &data}, &options.FindOneAndUpdateOptions{ReturnDocument: &after}).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Order{}, customerrors.ErrOrderNotFound
		}

		return models.Order{}, err
	}

	return updated, nil
}

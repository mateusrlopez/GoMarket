package repositories

import (
	"context"
	"time"

	"github.com/mateusrlopez/go-market/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentIntentsRepository interface {
	Create(intent models.PaymentIntent) error
	SetOneAsProcessingByGatewayID(gatewayId, paymentMethodType string) error
	SetOneAsSuccessfullByGatewayID(gatewayId string) error
	SetOneAsFailedByGatewayID(gatewayId string) error
}

type mongoPaymentIntentsRepository struct {
	mongo *mongo.Database
}

func NewPaymentIntentsRepository(mongo *mongo.Database) PaymentIntentsRepository {
	return mongoPaymentIntentsRepository{
		mongo: mongo,
	}
}

func (r mongoPaymentIntentsRepository) Create(intent models.PaymentIntent) error {
	intent.ID = primitive.NewObjectID().Hex()
	intent.CreatedAt = time.Now()

	_, err := r.mongo.Collection("payment_intents").InsertOne(context.Background(), &intent)

	if err != nil {
		return err
	}

	return nil
}

func (r mongoPaymentIntentsRepository) SetOneAsProcessingByGatewayID(gatewayId, paymentMethodType string) error {
	if err := r.mongo.Collection("payment_intents").FindOneAndUpdate(context.Background(), bson.M{"gatewayId": bson.M{"$eq": gatewayId}}, bson.M{"$set": bson.M{"status": "processing", "paymentMethodType": paymentMethodType}}).Err(); err != nil {
		return err
	}

	return nil
}

func (r mongoPaymentIntentsRepository) SetOneAsSuccessfullByGatewayID(gatewayId string) error {
	if err := r.mongo.Collection("payment_intents").FindOneAndUpdate(context.Background(), bson.M{"gatewayId": bson.M{"$eq": gatewayId}}, bson.M{"$set": bson.M{"status": "successfull", "succeededAt": time.Now()}}).Err(); err != nil {
		return err
	}

	return nil
}

func (r mongoPaymentIntentsRepository) SetOneAsFailedByGatewayID(gatewayId string) error {
	if err := r.mongo.Collection("payment_intents").FindOneAndUpdate(context.Background(), bson.M{"gatewayId": bson.M{"$eq": gatewayId}}, bson.M{"$set": bson.M{"status": "failed", "failedAt": time.Now()}}).Err(); err != nil {
		return err
	}

	return nil
}

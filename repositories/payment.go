package repositories

import (
	"context"

	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/shared/types"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepository struct {
	Collection   *mongo.Collection
	StripeClient *client.API
}

func (r PaymentRepository) Create(payment *models.Payment) (*mongo.InsertOneResult, error) {
	payment.BeforeInsert()

	return r.Collection.InsertOne(context.Background(), payment)
}

func (r PaymentRepository) StripeCreate(payment *models.Payment) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(int64(payment.Price.Value * 100)),
		Currency:      stripe.String(payment.Price.Currency),
		PaymentMethod: stripe.String(payment.Source),
		Confirm:       stripe.Bool(true),
	}

	return r.StripeClient.PaymentIntents.New(params)
}

func (r PaymentRepository) RetrieveAll(filter *types.PaymentIndexQuery) ([]models.Payment, error) {
	var payments []models.Payment

	cursor, err := r.Collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &payments); err != nil {
		return nil, err
	}

	return payments, nil
}

func (r PaymentRepository) RetrieveByID(id primitive.ObjectID, payment *models.Payment) error {
	return r.Collection.FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Decode(payment)
}

func (r PaymentRepository) Update(id primitive.ObjectID, payment *models.Payment) (*mongo.UpdateResult, error) {
	return r.Collection.UpdateOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": payment})
}

func (r PaymentRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return r.Collection.DeleteOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}})
}

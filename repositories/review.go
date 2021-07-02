package repositories

import (
	"context"

	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReviewRepository struct {
	Collection *mongo.Collection
}

func (r *ReviewRepository) Create(review *models.Review) (*mongo.InsertOneResult, error) {
	review.BeforeInsert()

	return r.Collection.InsertOne(context.Background(), review)
}

func (r *ReviewRepository) RetrieveAll(filter *types.ReviewIndexQuery) ([]models.Review, error) {
	var reviews []models.Review

	cursor, err := r.Collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *ReviewRepository) Update(id primitive.ObjectID, review *models.Review) (*mongo.UpdateResult, error) {
	return r.Collection.UpdateOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": review})
}

func (r *ReviewRepository) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return r.Collection.DeleteOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}})
}

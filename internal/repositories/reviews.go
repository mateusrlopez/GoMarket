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

type ReviewsRepository interface {
	Create(review models.Review) (models.Review, error)
	FindMany(filter models.Review) ([]models.Review, error)
	FindOneByID(id string) (models.Review, error)
	UpdateOneByID(id string, data models.Review) (models.Review, error)
	DeleteOneByID(id string) error
}

type mongoReviewsRepository struct {
	mongo *mongo.Database
}

func NewReviewsRepository(mongo *mongo.Database) ReviewsRepository {
	return mongoReviewsRepository{
		mongo: mongo,
	}
}

func (r mongoReviewsRepository) Create(review models.Review) (models.Review, error) {
	review.ID = primitive.NewObjectID().Hex()
	review.CreatedAt = time.Now()

	_, err := r.mongo.Collection("reviews").InsertOne(context.Background(), &review)

	if err != nil {
		return models.Review{}, err
	}

	return review, nil
}

func (r mongoReviewsRepository) FindMany(filter models.Review) ([]models.Review, error) {
	var reviews []models.Review

	cursor, err := r.mongo.Collection("reviews").Find(context.Background(), &filter)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var review models.Review

		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r mongoReviewsRepository) FindOneByID(id string) (models.Review, error) {
	var review models.Review

	if err := r.mongo.Collection("reviews").FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&review); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Review{}, customerrors.ErrReviewNotFound
		}

		return models.Review{}, err
	}

	return review, nil
}

func (r mongoReviewsRepository) UpdateOneByID(id string, data models.Review) (models.Review, error) {
	var updated models.Review

	after := options.After

	if err := r.mongo.Collection("reviews").FindOneAndUpdate(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": &data}, &options.FindOneAndUpdateOptions{ReturnDocument: &after}).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Review{}, customerrors.ErrReviewNotFound
		}

		return models.Review{}, err
	}

	return updated, nil
}

func (r mongoReviewsRepository) DeleteOneByID(id string) error {
	if err := r.mongo.Collection("reviews").FindOneAndDelete(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return customerrors.ErrReviewNotFound
		}

		return err
	}

	return nil
}

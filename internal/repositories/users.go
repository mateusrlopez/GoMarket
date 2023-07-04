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

type UsersRepository interface {
	Create(user models.User) (models.User, error)
	FindMany() ([]models.User, error)
	FindOneByID(id string) (models.User, error)
	FindOneByEmail(email string) (models.User, error)
	UpdateOneByID(id string, data models.User) (models.User, error)
	DeleteOneByID(id string) error
}

type mongoUserRepository struct {
	mongo *mongo.Database
}

func NewUsersRepository(mongo *mongo.Database) UsersRepository {
	return mongoUserRepository{
		mongo: mongo,
	}
}

func (r mongoUserRepository) Create(user models.User) (models.User, error) {
	user.ID = primitive.NewObjectID().Hex()
	user.CreatedAt = time.Now()

	_, err := r.mongo.Collection("users").InsertOne(context.Background(), &user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r mongoUserRepository) FindMany() ([]models.User, error) {
	var users []models.User

	cursor, err := r.mongo.Collection("users").Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var user models.User

		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r mongoUserRepository) FindOneByID(id string) (models.User, error) {
	var user models.User

	if err := r.mongo.Collection("users").FindOne(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, customerrors.ErrUserNotFound
		}

		return models.User{}, err
	}

	return user, nil
}

func (r mongoUserRepository) FindOneByEmail(email string) (models.User, error) {
	var user models.User

	if err := r.mongo.Collection("users").FindOne(context.Background(), bson.M{"email": bson.M{"$eq": email}}).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, customerrors.ErrUserNotFound
		}

		return models.User{}, err
	}

	return user, nil
}

func (r mongoUserRepository) UpdateOneByID(id string, data models.User) (models.User, error) {
	var updated models.User

	after := options.After

	if err := r.mongo.Collection("users").FindOneAndUpdate(context.Background(), bson.M{"_id": bson.M{"$eq": id}}, bson.M{"$set": &data}, &options.FindOneAndUpdateOptions{ReturnDocument: &after}).Decode(&updated); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, customerrors.ErrUserNotFound
		}

		return models.User{}, err
	}

	return updated, nil
}

func (r mongoUserRepository) DeleteOneByID(id string) error {
	if err := r.mongo.Collection("users").FindOneAndDelete(context.Background(), bson.M{"_id": bson.M{"$eq": id}}).Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return customerrors.ErrUserNotFound
		}

		return err
	}

	return nil
}

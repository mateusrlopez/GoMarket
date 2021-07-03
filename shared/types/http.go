package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReviewIndexQuery struct {
	ProductID primitive.ObjectID `schema:"productId" bson:"productId"`
}

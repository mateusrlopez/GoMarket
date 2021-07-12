package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReviewIndexQuery struct {
	ProductID primitive.ObjectID `schema:"productId" bson:"productId"`
}

type SkuIndexQuery struct {
	ProductID primitive.ObjectID `schema:"productId" bson:"productId"`
}

type OrderIndexQuery struct {
	UserID primitive.ObjectID `schema:"userId" bson:"userId"`
}

type PaymentIndexQuery struct {
	UserID primitive.ObjectID `schema:"userId" bson:"userId"`
}

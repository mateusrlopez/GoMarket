package utils

import (
	"reflect"

	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DecodeQuery(s interface{}, query map[string][]string) error {
	decoder := schema.NewDecoder()

	decoder.IgnoreUnknownKeys(true)

	baseId, _ := primitive.ObjectIDFromHex("")
	decoder.RegisterConverter(baseId, func(v string) reflect.Value {
		value, _ := primitive.ObjectIDFromHex(v)
		return reflect.ValueOf(value)
	})

	return decoder.Decode(s, query)
}

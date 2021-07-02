package utils

import (
	"reflect"

	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DecodeQuery(s interface{}, query map[string][]string) error {
	decoder := schema.NewDecoder()
	baseId, _ := primitive.ObjectIDFromHex("")

	decoder.IgnoreUnknownKeys(true)
	decoder.RegisterConverter(baseId, func(v string) reflect.Value {
		value, _ := primitive.ObjectIDFromHex(v)
		return reflect.ValueOf(value)
	})

	return decoder.Decode(s, query)
}

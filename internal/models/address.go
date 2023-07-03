package models

type Address struct {
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	Country    string `json:"country" bson:"country"`
	PostalCode string `json:"postalCode" bson:"postalCode"`
	Line1      string `json:"line1" bson:"line1"`
	Line2      string `json:"line2" bson:"line2"`
}

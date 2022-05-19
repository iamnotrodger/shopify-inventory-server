package model

type Location struct {
	Street   string `json:"street,omitempty" bson:"street,omitempty"`
	City     string `json:"city,omitempty" bson:"city,omitempty"`
	Province string `json:"province,omitempty" bson:"province,omitempty"`
	Country  string `json:"country,omitempty" bson:"country,omitempty"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Warehouse struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Location    *Location          `json:"location,omitempty" bson:"location,omitempty"`
	Inventories []*Inventory       `json:"inventories,omitempty" bson:"inventories,omitempty"`
}

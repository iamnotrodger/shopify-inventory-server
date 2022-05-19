package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Price      float64            `json:"price,omitempty" bson:"price,omitempty"`
	Warehouses []*Warehouse       `json:"warehouses,omitempty" bson:"warehouses,omitempty"`
}

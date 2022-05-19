package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Price      float64            `json:"price,omitempty" bson:"price,omitempty"`
	Warehouses []*Warehouse       `json:"warehouses,omitempty" bson:"warehouses,omitempty"`
}

func (i *Inventory) Validate() error {
	if i.Name == "" {
		return &ErrInvalidModel{
			message: "inventory requires a name",
		}
	}

	if i.Price < 0 {
		return &ErrInvalidModel{
			message: "inventory price can't be less than zero",
		}
	}

	return nil
}

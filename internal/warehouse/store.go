package warehouse

import (
	"context"
	"fmt"

	"github.com/iamnotrodger/shopify-inventory-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	INVENTORY = "inventory"
	WAREHOUSE = "warehouse"
)

type Store struct {
	db *mongo.Database
}

func NewStore(db *mongo.Database) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Find(ctx context.Context, warehouseID string) (*model.Warehouse, error) {
	id, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return nil, primitive.ErrInvalidHex
	}

	singleRes := s.db.Collection(WAREHOUSE).FindOne(ctx, bson.M{"_id": id}, &options.FindOneOptions{})
	if err = singleRes.Err(); err != nil {
		return nil, err
	}

	warehouse := &model.Warehouse{}
	err = singleRes.Decode(warehouse)
	if err != nil {
		err = fmt.Errorf("error decoding inventory: %w", err)
		return nil, err
	}

	return warehouse, nil
}

func (s *Store) FindInventories(ctx context.Context, warehouseID string) ([]*model.Inventory, error) {
	return nil, nil
}

func (s *Store) Insert(ctx context.Context, inventory *model.Inventory) error {
	return nil
}

// InsertInventory is a transaction that adds inventory to warehouse and add warehouse to inventory
func (s *Store) InsertInventory(ctx context.Context, warehouseID string, inventoryID string) error {
	return nil
}

func (s *Store) Delete(ctx context.Context, inventoryID string) error {
	return nil
}

// DeleteInventory is a transaction delete inventory to warehouse and delete warehouse to inventory
func (s *Store) DeleteInventory(ctx context.Context, warehouseID string, inventoryID string) error {
	return nil
}

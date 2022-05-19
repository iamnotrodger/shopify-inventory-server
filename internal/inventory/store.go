package inventory

import (
	"context"
	"fmt"

	"github.com/iamnotrodger/shopify-inventory-server/internal/model"
	"github.com/iamnotrodger/shopify-inventory-server/internal/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	collection *mongo.Collection
}

func NewStore(db *mongo.Database) *Store {
	return &Store{
		collection: db.Collection("inventory"),
	}
}

func (s *Store) Find(ctx context.Context, inventoryID string) (*model.Inventory, error) {
	id, err := primitive.ObjectIDFromHex(inventoryID)
	if err != nil {
		return nil, primitive.ErrInvalidHex
	}

	singleRes := s.collection.FindOne(ctx, bson.M{"_id": id}, &options.FindOneOptions{})
	if err = singleRes.Err(); err != nil {
		return nil, err
	}

	inventory := &model.Inventory{}
	err = singleRes.Decode(inventory)
	if err != nil {
		err = fmt.Errorf("error decoding inventory: %w", err)
		return nil, err
	}

	return inventory, nil
}

func (s *Store) FindWarehouses(ctx context.Context, inventoryID string, queryParam ...query.QueryParams) ([]*model.Warehouse, error) {
	id, err := primitive.ObjectIDFromHex(inventoryID)
	if err != nil {
		return nil, primitive.ErrInvalidHex
	}

	match := bson.D{{Key: "$match", Value: bson.M{"_id": id}}}
	matchWarehouse := bson.D{{
		Key: "$match",
		Value: bson.D{{
			Key: "$expr",
			Value: bson.D{{
				Key:   "$in",
				Value: bson.A{"$_id", "$$warehouse_ids"},
			}},
		}},
	}}

	lookupPipeline := bson.A{matchWarehouse}
	if len(queryParam) > 0 {
		for _, queryOpts := range queryParam[0].GetPipeline() {
			lookupPipeline = append(lookupPipeline, queryOpts)
		}
	}

	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: "warehouse"},
			{Key: "let", Value: bson.D{{Key: "warehouse_ids", Value: "$warehouse_ids"}}},
			{Key: "pipeline", Value: lookupPipeline},
			{Key: "as", Value: "warehouses"},
		},
	}}

	pipeline := mongo.Pipeline{match, lookup}
	cursor, err := s.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if cursor.RemainingBatchLength() < 1 {
		return nil, mongo.ErrNoDocuments
	}

	inventory := model.Inventory{}
	cursor.Next(ctx)
	cursor.Decode(&inventory)
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return inventory.Warehouses, nil
}

func (s *Store) FindMany(ctx context.Context, queryParam ...query.QueryParams) ([]*model.Inventory, error) {
	return nil, nil
}

func (s *Store) Insert(ctx context.Context, inventory *model.Inventory) error {
	return nil
}

func (s *Store) Delete(ctx context.Context, inventoryID string) error {
	return nil
}

func (s *Store) Update(ctx context.Context, inventoryID string, inventory *model.Inventory) error {
	return nil
}

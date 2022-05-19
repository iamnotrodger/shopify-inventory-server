package warehouse

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

func (s *Store) FindInventories(ctx context.Context, warehouseID string, queryParam ...query.QueryParams) ([]*model.Inventory, error) {
	id, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return nil, primitive.ErrInvalidHex
	}

	match := bson.D{{Key: "$match", Value: bson.M{"_id": id}}}
	matchArtists := bson.D{{
		Key: "$match",
		Value: bson.D{{
			Key: "$expr",
			Value: bson.D{{
				Key:   "$in",
				Value: bson.A{"$_id", "$$inventory_ids"},
			}},
		}},
	}}

	lookupPipeline := bson.A{matchArtists}
	if len(queryParam) > 0 {
		for _, queryOpts := range queryParam[0].GetPipeline() {
			lookupPipeline = append(lookupPipeline, queryOpts)
		}
	}

	lookup := bson.D{{
		Key: "$lookup",
		Value: bson.D{
			{Key: "from", Value: "inventory"},
			{Key: "let", Value: bson.D{{Key: "inventory_ids", Value: "$inventory_ids"}}},
			{Key: "pipeline", Value: lookupPipeline},
			{Key: "as", Value: "inventories"},
		},
	}}

	pipeline := mongo.Pipeline{match, lookup}
	cursor, err := s.db.Collection(WAREHOUSE).Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if cursor.RemainingBatchLength() < 1 {
		return nil, mongo.ErrNoDocuments
	}

	warehouse := model.Warehouse{}
	cursor.Next(ctx)
	cursor.Decode(&warehouse)
	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return warehouse.Inventories, nil
}

func (s *Store) Insert(ctx context.Context, warehouse *model.Warehouse) error {
	warehouse.ID = primitive.NewObjectID()
	_, err := s.db.Collection(WAREHOUSE).InsertOne(ctx, warehouse)
	return err
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

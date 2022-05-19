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
	INVENTORY       = "inventory"
	WAREHOUSE       = "warehouse"
	INVENTORY_FIELD = "inventory_ids"
	WAREHOUSE_FIELD = "warehouse_ids"
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

func (s *Store) FindMany(ctx context.Context, queryParam ...query.QueryParams) ([]*model.Warehouse, error) {
	var opts *options.FindOptions
	filter := bson.D{}

	if len(queryParam) > 0 {
		filter = queryParam[0].GetFilter()
		opts = queryParam[0].GetFindOptions()
	}

	cursor, err := s.db.Collection(WAREHOUSE).Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	warehouses := []*model.Warehouse{}
	err = cursor.All(ctx, &warehouses)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal warehouse: %w", err)
		return nil, err
	}

	return warehouses, nil
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
	warehouseIDObj, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return primitive.ErrInvalidHex
	}
	inventoryIDObj, err := primitive.ObjectIDFromHex(inventoryID)
	if err != nil {
		return primitive.ErrInvalidHex
	}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Add inventory to warehouse
		err := s.appendElement(sessCtx, &listOperationConfig{
			collection: WAREHOUSE,
			field:      WAREHOUSE_FIELD,
			id:         warehouseIDObj,
			element:    inventoryIDObj,
		})
		if err != nil {
			return nil, err
		}

		// Add warehouse to inventory
		err = s.appendElement(sessCtx, &listOperationConfig{
			collection: INVENTORY,
			field:      INVENTORY_FIELD,
			id:         warehouseIDObj,
			element:    inventoryIDObj,
		})
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	session, err := s.db.Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Delete(ctx context.Context, warehouseID string) error {
	id, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return err
	}

	res, err := s.db.Collection(WAREHOUSE).DeleteOne(ctx, bson.M{"_id": id})
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return err
}

// DeleteInventory is a transaction delete inventory to warehouse and delete warehouse to inventory
func (s *Store) DeleteInventory(ctx context.Context, warehouseID string, inventoryID string) error {
	warehouseIDObj, err := primitive.ObjectIDFromHex(warehouseID)
	if err != nil {
		return primitive.ErrInvalidHex
	}
	inventoryIDObj, err := primitive.ObjectIDFromHex(inventoryID)
	if err != nil {
		return primitive.ErrInvalidHex
	}

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Remove inventory to warehouse
		err := s.removeElement(sessCtx, &listOperationConfig{
			collection: WAREHOUSE,
			field:      WAREHOUSE_FIELD,
			id:         warehouseIDObj,
			element:    inventoryIDObj,
		})
		if err != nil {
			return nil, err
		}

		// Remove warehouse to inventory
		err = s.removeElement(sessCtx, &listOperationConfig{
			collection: INVENTORY,
			field:      INVENTORY_FIELD,
			id:         warehouseIDObj,
			element:    inventoryIDObj,
		})
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	session, err := s.db.Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		return err
	}

	return nil
}

type listOperationConfig struct {
	collection string
	field      string
	id         primitive.ObjectID
	element    interface{}
}

func (s *Store) appendElement(ctx context.Context, config *listOperationConfig) error {
	res, err := s.db.Collection(config.collection).UpdateByID(ctx, config.id, bson.M{
		"$push": bson.M{config.field: config.element},
	})
	if err != nil {
		return err
	} else if res.MatchedCount == 0 {
		return mongo.ErrNilDocument
	}
	return nil
}

func (s *Store) removeElement(ctx context.Context, config *listOperationConfig) error {
	res, err := s.db.Collection(config.collection).UpdateByID(ctx, config.id, bson.M{
		"$pull": bson.M{config.field: config.element},
	})
	if err != nil {
		return err
	} else if res.MatchedCount == 0 {
		return mongo.ErrNilDocument
	}
	return nil
}

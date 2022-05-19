package query

import (
	"strconv"

	"github.com/iamnotrodger/shopify-inventory-server/cmd/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InventoryQueryParams struct {
	limit int64
	skip  int64
	sort  map[string]int
}

func NewInventoryQuery(parameters map[string][]string) *InventoryQueryParams {
	query := &InventoryQueryParams{}

	if limit, ok := parameters["limit"]; ok {
		query.setLimitFromString(limit[0])
	} else {
		query.limit = config.Global.InventoryLimit
	}
	if skip, ok := parameters["skip"]; ok {
		query.setSkipFromString(skip[0])
	}
	if sort, ok := parameters["sort"]; ok {
		query.SetSort(sort)
	}

	return query
}

func (q *InventoryQueryParams) GetFilter() bson.D {
	return bson.D{}
}

func (q *InventoryQueryParams) GetFindOptions() *options.FindOptions {
	options := options.Find()
	if q.isSortValid() {
		sort := getSortAsBson(q.sort)
		options.SetSort(sort)
	}
	if q.isSkipValid() {
		options.SetSkip(q.skip)
	}
	options.SetLimit(q.limit)
	return options
}

func (q *InventoryQueryParams) GetPipeline() []bson.D {
	pipeline := []bson.D{}

	if q.isSortValid() {
		sort := bson.D{{Key: "$sort", Value: q.sort}}
		pipeline = append(pipeline, sort)
	}
	if q.isSkipValid() {
		skip := bson.D{{Key: "$skip", Value: q.skip}}
		pipeline = append(pipeline, skip)
	}
	limit := bson.D{{Key: "$limit", Value: q.limit}}
	pipeline = append(pipeline, limit)

	return pipeline
}

func (q *InventoryQueryParams) SetLimit(limit int64) {
	if limit < config.Global.InventoryLimitMin {
		q.limit = config.Global.InventoryLimit
	} else if limit > config.Global.InventoryLimitMax {
		q.limit = config.Global.InventoryLimitMax
	} else {
		q.limit = limit
	}
}
func (q *InventoryQueryParams) SetSkip(skip int64) {
	if skip > 0 {
		q.skip = skip
	}
}

func (q *InventoryQueryParams) SetSort(sortArray []string) {
	q.sort = map[string]int{}
	for _, sortString := range sortArray {
		key, value := parseSort(sortString)
		if key != "" {
			q.sort[key] = value
		}
	}
}

func (q *InventoryQueryParams) setLimitFromString(limitString string) {
	limit, err := strconv.ParseInt(limitString, 0, 64)
	if err != nil {
		q.limit = config.Global.InventoryLimit
	} else {
		q.SetLimit(limit)
	}
}

func (q *InventoryQueryParams) setSkipFromString(skipString string) {
	skip, err := strconv.ParseInt(skipString, 0, 64)
	if err == nil {
		q.SetSkip(skip)
	}
}

func (q *InventoryQueryParams) isSortValid() bool {
	return q.sort != nil && len(q.sort) > 0
}

func (q *InventoryQueryParams) isSkipValid() bool {
	return q.skip > 0
}

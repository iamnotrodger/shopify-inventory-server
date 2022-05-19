package config

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const (
	defaultPort        = 8080
	defaultMongoURI    = "mongodb://localhost:27017"
	defaultMongoDBName = "shopify-inventory"

	defaultInventoryLimit    = 15
	defaultInventoryLimitMin = 1
	defaultInventoryLimitMax = 100

	defaultWarehouseLimit    = 15
	defaultWarehouseLimitMin = 1
	defaultWarehouseLimitMax = 100
)

type Spec struct {
	Port              int    `mapstructure:"port"`
	MongoURI          string `mapstructure:"mongo_uri"`
	MongoDBName       string `mapstructure:"mongo_db_name"`
	InventoryLimit    int64  `mapstructure:"inventory_limit"`
	InventoryLimitMin int64  `mapstructure:"inventory_limit_min"`
	InventoryLimitMax int64  `mapstructure:"inventory_limit_max"`
	WarehouseLimit    int64  `mapstructure:"warehouse_limit"`
	WarehouseLimitMin int64  `mapstructure:"warehouse_limit_min"`
	WarehouseLimitMax int64  `mapstructure:"warehouse_limit_max"`
}

var Global = Spec{
	Port:              defaultPort,
	MongoURI:          defaultMongoURI,
	MongoDBName:       defaultMongoDBName,
	InventoryLimit:    defaultInventoryLimit,
	InventoryLimitMin: defaultInventoryLimitMin,
	InventoryLimitMax: defaultInventoryLimitMax,
	WarehouseLimit:    defaultWarehouseLimit,
	WarehouseLimitMin: defaultWarehouseLimitMin,
	WarehouseLimitMax: defaultWarehouseLimitMax,
}

func LoadConfig() {
	v := viper.New()
	v.SetConfigFile(".env")
	v.ReadInConfig()
	v.AutomaticEnv()

	setDefaults(v, Global)

	if err := v.Unmarshal(&Global); err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config %s", err))
	}
}

func setDefaults(v *viper.Viper, i interface{}) {
	values := map[string]interface{}{}
	if err := mapstructure.Decode(i, &values); err != nil {
		panic(err)
	}
	for key, defaultValue := range values {
		v.SetDefault(key, defaultValue)
	}
}

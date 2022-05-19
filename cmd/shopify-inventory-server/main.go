package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/iamnotrodger/shopify-inventory-server/cmd/config"
	"github.com/iamnotrodger/shopify-inventory-server/internal/inventory"
	"github.com/iamnotrodger/shopify-inventory-server/internal/middleware"
	"github.com/iamnotrodger/shopify-inventory-server/internal/util"
	"github.com/iamnotrodger/shopify-inventory-server/internal/warehouse"
	"github.com/rs/cors"
)

func main() {
	config.LoadConfig()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := util.GetMongoClient(ctx, config.Global.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database(config.Global.MongoDBName)

	inventoryHandler := inventory.NewHandler(db)
	warehouseHandler := warehouse.NewHandler(db)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggingMiddleware)

	// Inventory Routes
	inventoryHandler.RegisterRoutes(router)
	// Warehouse Routes
	warehouseHandler.RegisterRoutes(router)

	server := &http.Server{
		Handler:      cors.Default().Handler(router),
		Addr:         fmt.Sprintf("127.0.0.1:%v", config.Global.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("API Started. Listening on", config.Global.Port)
	log.Fatal(server.ListenAndServe())
}

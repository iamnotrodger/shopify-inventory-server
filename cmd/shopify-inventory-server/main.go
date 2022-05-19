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

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.LoggingMiddleware)

	//Inventory Routes
	inventoryHandler.RegisterRoutes(router)

	server := cors.Default().Handler(router)
	log.Println("API Started. Listening on", config.Global.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", config.Global.Port), server))
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/iamnotrodger/shopify-inventory-server/cmd/config"
	"github.com/iamnotrodger/shopify-inventory-server/internal/inventory"
	"github.com/iamnotrodger/shopify-inventory-server/internal/middleware"
	"github.com/iamnotrodger/shopify-inventory-server/internal/util"
	"github.com/iamnotrodger/shopify-inventory-server/internal/warehouse"
	"github.com/rs/cors"
)

var flagDev = flag.Bool("dev", true, "Run in development mode")

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

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
	// SPA Routes
	if *flagDev {
		proxyUrl, err := url.Parse(config.Global.ProxyRoute)
		if err != nil {
			log.Fatal(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
		router.PathPrefix("/").Handler(proxy)
	} else {
		spa := spaHandler{staticPath: config.Global.StaticPath, indexPath: "index.html"}
		router.PathPrefix("/").Handler(spa)
	}

	server := &http.Server{
		Handler:      cors.Default().Handler(router),
		Addr:         fmt.Sprintf("127.0.0.1:%v", config.Global.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("API Started. Listening on", config.Global.Port)
	log.Fatal(server.ListenAndServe())
}

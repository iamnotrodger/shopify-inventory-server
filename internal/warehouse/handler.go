package warehouse

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamnotrodger/shopify-inventory-server/internal/model"
	"github.com/iamnotrodger/shopify-inventory-server/internal/query"
	"github.com/iamnotrodger/shopify-inventory-server/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	store *Store
}

func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		store: NewStore(db),
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/warehouse", h.Post).Methods("POST")
	router.HandleFunc("/api/warehouse", h.GetMany).Methods("GET")
	router.HandleFunc("/api/warehouse/{id}", h.Get).Methods("GET")
	router.HandleFunc("/api/warehouse/{id}/inventory", h.GetInventories).Methods("GET")
	router.HandleFunc("/api/warehouse/{id}/inventory/{inventoryID}", h.PostInventory).Methods("POST")
	router.HandleFunc("/api/warehouse/{id}/inventory/{inventoryID}", h.DeleteInventory).Methods("DELETE")
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	warehouseID := params["id"]

	warehouse, err := h.store.Find(r.Context(), warehouseID)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(warehouse)
}

func (h *Handler) GetMany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryParam := query.NewInventoryQuery(r.URL.Query())
	warehouses, err := h.store.FindMany(r.Context(), queryParam)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(warehouses)
}

func (h *Handler) GetInventories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	warehouseID := params["id"]

	queryParam := query.NewWarehouseQuery(r.URL.Query())
	inventories, err := h.store.FindInventories(r.Context(), warehouseID, queryParam)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(inventories)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var warehouse model.Warehouse
	err := json.NewDecoder(r.Body).Decode(&warehouse)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	err = warehouse.Validate()
	if err != nil {
		util.HandleError(w, err)
		return
	}

	err = h.store.Insert(r.Context(), &warehouse)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(&warehouse)
}

// PostInventory will add inventory to warehouse and add warehouse to the inventory
func (h *Handler) PostInventory(w http.ResponseWriter, r *http.Request) {
	// TODO: validate if warehouse exists
	// TODO: validate if inventory exists
	// TODO: add inventory to warehouse
	// TODO: add warehouse to inventory
}

// PostInventory will delete inventory to warehouse and delete warehouse to the inventory
func (h *Handler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
	// TODO: validate if warehouse exists
	// TODO: validate if inventory exists
	// TODO: delete inventory to warehouse
	// TODO: delete warehouse to inventory
}

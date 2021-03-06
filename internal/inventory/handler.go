package inventory

import (
	"encoding/json"
	"fmt"
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
	router.HandleFunc("/api/inventory", h.Post).Methods("POST")
	router.HandleFunc("/api/inventory", h.GetMany).Methods("GET")
	router.HandleFunc("/api/inventory/{id}", h.Get).Methods("GET")
	router.HandleFunc("/api/inventory/{id}", h.Delete).Methods("DELETE")
	router.HandleFunc("/api/inventory/{id}", h.Update).Methods("PUT")
	router.HandleFunc("/api/inventory/{id}", h.Patch).Methods("PATCH")
	router.HandleFunc("/api/inventory/{id}/warehouse", h.GetWarehouses).Methods("GET")
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	inventoryID := params["id"]

	inventory, err := h.store.Find(r.Context(), inventoryID)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(inventory)
}

func (h *Handler) GetWarehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	inventoryID := params["id"]

	queryParams := query.NewWarehouseQuery(r.URL.Query())
	warehouse, err := h.store.FindWarehouses(r.Context(), inventoryID, queryParams)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(warehouse)
}

func (h *Handler) GetMany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryParam := query.NewInventoryQuery(r.URL.Query())
	inventories, err := h.store.FindMany(r.Context(), queryParam)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(inventories)
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var inventory model.Inventory
	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	err = inventory.Validate()
	if err != nil {
		util.HandleError(w, err)
		return
	}

	err = h.store.Insert(r.Context(), &inventory)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(&inventory)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inventoryID := params["id"]

	err := h.store.Delete(r.Context(), inventoryID)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	msg := fmt.Sprintf("inventory %v deleted", inventoryID)
	w.WriteHeader(200)
	w.Write([]byte(msg))
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	inventoryID := params["id"]

	var inventory model.Inventory
	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	if err = inventory.Validate(); err != nil {
		util.HandleError(w, err)
		return
	}

	err = h.store.Update(r.Context(), inventoryID, &inventory)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(&inventory)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	inventoryID := params["id"]

	var inventory model.Inventory
	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	err = h.store.Update(r.Context(), inventoryID, &inventory)
	if err != nil {
		util.HandleError(w, err)
		return
	}

	json.NewEncoder(w).Encode(&inventory)
}

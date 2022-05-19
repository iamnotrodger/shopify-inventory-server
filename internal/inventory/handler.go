package inventory

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
}

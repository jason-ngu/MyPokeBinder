package handler

import (
	"backend/internal/collections"
	"backend/internal/models"
	"backend/pkg/utilities"
	"encoding/json"
	"net/http"
	"strconv"
)

type collectionsHandler struct {
	collectionsService collections.Service
}

func NewCollectionsHandler(collectionsService collections.Service) collections.Handler {
	return &collectionsHandler{collectionsService: collectionsService}
}

func (h *collectionsHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var newCollection models.CollectionModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newCollection)
	if err != nil {
		http.Error(w, "Invalid Collection", http.StatusBadRequest)
		return
	}

	createdCollection, err := h.collectionsService.Create(ctx, &newCollection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(createdCollection)
}

func (h *collectionsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionIDStr := r.PathValue("collectionID")
	if collectionIDStr == "" {
		http.Error(w, "No Collection ID provided", http.StatusBadRequest)
		return
	}
	collectionID, err := strconv.Atoi(collectionIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection, err := h.collectionsService.GetByID(ctx, collectionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(collection)
}

func (h *collectionsHandler) GetAllCollectionsByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userIDStr := r.PathValue("userID")
	if userIDStr == "" {
		http.Error(w, "No User ID provided", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	pagQuery := ctx.Value(utilities.PaginationQuery{}).(*utilities.PaginationQuery)

	collectionList, err := h.collectionsService.GetAllCollectionsByUserID(ctx, userID, pagQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(collectionList)
}

func (h *collectionsHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var collectionToUpdate models.CollectionModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&collectionToUpdate)
	if err != nil {
		http.Error(w, "Invalid Collection", http.StatusBadRequest)
		return
	}

	updatedCollection, err := h.collectionsService.Update(ctx, &collectionToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(updatedCollection)
}

func (h *collectionsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionIDStr := r.PathValue("collectionID")
	if collectionIDStr == "" {
		http.Error(w, "No Collection ID provided", http.StatusBadRequest)
		return
	}

	collectionID, err := strconv.Atoi(collectionIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.collectionsService.Delete(ctx, collectionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

package handler

import (
	"backend/internal/collectioncards"
	"backend/internal/models"
	"backend/pkg/utilities"
	"encoding/json"
	"net/http"
	"strconv"
)

type collectioncardsHandler struct {
	collectioncardsService collectioncards.Service
}

func NewCollectionCardsHandler(collectioncardsService collectioncards.Service) collectioncards.Handler {
	return &collectioncardsHandler{collectioncardsService: collectioncardsService}
}

func (h *collectioncardsHandler) AddCardsToCollection(w http.ResponseWriter, r *http.Request) {
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

	var collectioncardsToAdd []*models.CollectionCardsModel
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&collectioncardsToAdd)
	if err != nil {
		http.Error(w, "Invalid Collection Cards", http.StatusBadRequest)
		return
	}

	collectioncards, err := h.collectioncardsService.AddCardsToCollection(ctx, collectionID, collectioncardsToAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(collectioncards)

}

func (h *collectioncardsHandler) GetByCollectionID(w http.ResponseWriter, r *http.Request) {
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

	pagQuery := ctx.Value(utilities.PaginationQuery{}).(*utilities.PaginationQuery)

	collectioncards, err := h.collectioncardsService.GetByCollectionID(ctx, collectionID, pagQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(collectioncards)
}

func (h *collectioncardsHandler) RemoveCardsFromCollection(w http.ResponseWriter, r *http.Request) {
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

	var collectioncardsToRemove []*models.CollectionCardsModel
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&collectioncardsToRemove)
	if err != nil {
		http.Error(w, "Invalid Collection Cards", http.StatusBadRequest)
		return
	}

	collectioncards, err := h.collectioncardsService.RemoveCardsFromCollection(ctx, collectionID, collectioncardsToRemove)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(collectioncards)
}

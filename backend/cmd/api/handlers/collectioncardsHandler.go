package handlers

import (
	"backend/internal/models"
	collectioncardsService "backend/internal/services/collectioncards"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetCollectionCards(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collectioncardId, err := strconv.Atoi(params["collectionId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	collectioncards, err := collectioncardsService.GetCollectionCards(h.ENV, collectioncardId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(collectioncards)
}

func (h handler) AddCardsToCollection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collectioncardId, err := strconv.Atoi(params["collectionId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var collectioncardsToAdd []models.CollectionCardsModel
	err = json.NewDecoder(r.Body).Decode(&collectioncardsToAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedCollectioncards, err := collectioncardsService.AddCardsToCollection(h.ENV, collectioncardId, collectioncardsToAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedCollectioncards)
}

func (h handler) RemoveCardsFromCollection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collectioncardId, err := strconv.Atoi(params["collectionId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var collectioncardsToAdd []models.CollectionCardsModel
	err = json.NewDecoder(r.Body).Decode(&collectioncardsToAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedCollectioncards, err := collectioncardsService.RemoveCardsFromCollection(h.ENV, collectioncardId, collectioncardsToAdd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedCollectioncards)
}

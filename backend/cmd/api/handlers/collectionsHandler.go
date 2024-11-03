package handlers

import (
	"backend/internal/models"
	collectionsService "backend/internal/services/collections"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetCollection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collectionId, err := strconv.Atoi(params["collectionId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	collection, err := collectionsService.GetCollection(h.ENV, collectionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(collection)
}

func (h handler) CreateCollection(w http.ResponseWriter, r *http.Request) {
	var collectionToCreate models.CollectionModel
	err := json.NewDecoder(r.Body).Decode(&collectionToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newCollection, err := collectionsService.CreateCollection(h.ENV, collectionToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newCollection)
}

func (h handler) UpdateCollection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collectionId, err := strconv.Atoi(params["collectionId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var collectionToUpdate models.CollectionModel
	err = json.NewDecoder(r.Body).Decode(&collectionToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updatedCollection, err := collectionsService.UpdateCollection(h.ENV, collectionId, collectionToUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedCollection)
}

func (h handler) DeleteCollection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	collectionId, err := strconv.Atoi(params["collectionId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	didDeleteCollection, err := collectionsService.DeleteCollection(h.ENV, collectionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(didDeleteCollection)
}

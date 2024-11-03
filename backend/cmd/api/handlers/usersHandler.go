package handlers

import (
	"backend/internal/models"
	collectionsService "backend/internal/services/collections"
	usersService "backend/internal/services/users"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := usersService.GetUser(h.ENV, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userToCreate models.UserModel
	err := json.NewDecoder(r.Body).Decode(&userToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdUser, err := usersService.CreateUser(h.ENV, userToCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdUser)
}

func (h handler) GetUserCollections(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	collections, err := collectionsService.GetUserCollections(h.ENV, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(collections)
}

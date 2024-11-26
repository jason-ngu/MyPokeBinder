package handler

import (
	"backend/internal/models"
	"backend/internal/users"
	"encoding/json"
	"net/http"
)

type usersHandler struct {
	usersService users.Service
}

func NewUsersHandler(usersService users.Service) users.Handler {
	return &usersHandler{usersService: usersService}
}

func (h *usersHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var newUser models.UserModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid User", http.StatusBadRequest)
		return
	}

	createdUser, err := h.usersService.Create(ctx, &newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(createdUser)
}

// func (h *usersHandler) GetByID(w http.ResponseWriter, r *http.Request) {

// }

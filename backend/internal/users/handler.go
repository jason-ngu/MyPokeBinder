package users

import "net/http"

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	// GetByID(w http.ResponseWriter, r *http.Request)
}

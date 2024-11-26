package cards

import "net/http"

type Handler interface {
	GetByID(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
}

package collectioncards

import "net/http"

type Handler interface {
	AddCardsToCollection(w http.ResponseWriter, r *http.Request)
	GetByCollectionID(w http.ResponseWriter, r *http.Request)
	RemoveCardsFromCollection(w http.ResponseWriter, r *http.Request)
}

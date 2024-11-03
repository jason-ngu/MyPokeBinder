package handlers

import (
	"backend/internal/models"
	cardsService "backend/internal/services/cards"
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func (h handler) SearchCards(w http.ResponseWriter, r *http.Request) {
	var cardSearchParams models.CardSearchParams
	err := decoder.Decode(&cardSearchParams, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cards, err := cardsService.SearchCards(h.ENV, cardSearchParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cards)
}

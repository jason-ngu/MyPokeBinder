package handler

import (
	"backend/internal/cards"
	"backend/internal/models"
	"backend/pkg/utilities"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"
)

type cardsHandler struct {
	cardsService cards.Service
}

func NewCardsHandler(cardsService cards.Service) cards.Handler {
	return &cardsHandler{cardsService: cardsService}
}

var decoder = schema.NewDecoder()

func (h *cardsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cardIDStr := r.PathValue("cardID")
	if cardIDStr == "" {
		http.Error(w, "No Card ID provided", http.StatusBadRequest)
		return
	}
	cardID, err := strconv.Atoi(cardIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	card, err := h.cardsService.GetByID(ctx, cardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(card)
}

func (h *cardsHandler) Search(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var cardSearchParams models.CardSearchParams
	err := decoder.Decode(&cardSearchParams, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pagQuery := ctx.Value(utilities.PaginationQuery{}).(*utilities.PaginationQuery)

	cards, err := h.cardsService.Search(ctx, &cardSearchParams, pagQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cards)
}

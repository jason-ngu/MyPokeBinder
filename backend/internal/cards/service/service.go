package service

import (
	"backend/internal/cards"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type cardsService struct {
	cardsRepo cards.Repository
}

func NewCardsService(cardsRepo cards.Repository) cards.Service {
	return &cardsService{cardsRepo: cardsRepo}
}

func (s *cardsService) Create(ctx context.Context, newCard *models.CardModel) (*models.CardModel, error) {
	return s.cardsRepo.Create(ctx, newCard)
}

func (s *cardsService) GetByID(ctx context.Context, cardID int) (*models.CardModel, error) {
	return s.cardsRepo.GetByID(ctx, cardID)
}

func (s *cardsService) Search(ctx context.Context, searchParams *models.CardSearchParams, query *utilities.PaginationQuery) (*models.CardsList, error) {
	return s.cardsRepo.Search(ctx, searchParams, query)
}

func (s *cardsService) Update(ctx context.Context, cardID int, cardToUpdate *models.CardModel) (*models.CardModel, error) {
	return s.cardsRepo.Update(ctx, cardID, cardToUpdate)
}

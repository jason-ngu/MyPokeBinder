package service

import (
	"backend/internal/cardtypes"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type cardtypesService struct {
	cardtypesRepo cardtypes.Repository
}

func NewCardTypesService(cardtypesRepo cardtypes.Repository) cardtypes.Service {
	return &cardtypesService{cardtypesRepo: cardtypesRepo}
}

func (s *cardtypesService) Create(ctx context.Context, newCardType *models.CardTypeEntity) (*models.CardTypeEntity, error) {
	return s.cardtypesRepo.Create(ctx, newCardType)
}
func (s *cardtypesService) GetByCardID(ctx context.Context, cardID int, query *utilities.PaginationQuery) (*models.CardTypesList, error) {
	return s.cardtypesRepo.GetByCardID(ctx, cardID, query)
}

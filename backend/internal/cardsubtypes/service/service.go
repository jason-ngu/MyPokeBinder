package service

import (
	"backend/internal/cardsubtypes"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type cardsubtypesService struct {
	cardsubtypesRepo cardsubtypes.Repository
}

func NewCardSubtypesService(cardsubtypesRepo cardsubtypes.Repository) cardsubtypes.Service {
	return &cardsubtypesService{cardsubtypesRepo: cardsubtypesRepo}
}

func (s *cardsubtypesService) Create(ctx context.Context, newCardSubtype *models.CardSubtypeEntity) (*models.CardSubtypeEntity, error) {
	return s.cardsubtypesRepo.Create(ctx, newCardSubtype)
}
func (s *cardsubtypesService) GetByCardID(ctx context.Context, cardID int, query *utilities.PaginationQuery) (*models.CardSubtypesList, error) {
	return s.cardsubtypesRepo.GetByCardID(ctx, cardID, query)
}

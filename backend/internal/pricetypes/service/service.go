package service

import (
	"backend/internal/models"
	"backend/internal/pricetypes"
	"backend/pkg/utilities"
	"context"
)

type pricetypesService struct {
	pricetypesRepo pricetypes.Repository
}

func NewPricetypesService(pricetypesRepo pricetypes.Repository) pricetypes.Service {
	return &pricetypesService{pricetypesRepo: pricetypesRepo}
}

func (s *pricetypesService) Create(ctx context.Context, newPricetype *models.PricetypeEntity) (*models.PricetypeEntity, error) {
	return s.pricetypesRepo.Create(ctx, newPricetype)
}

func (s *pricetypesService) GetByID(ctx context.Context, pricetypeID int) (*models.PricetypeModel, error) {
	return s.pricetypesRepo.GetByID(ctx, pricetypeID)
}

func (s *pricetypesService) GetAllPricetypes(ctx context.Context, query *utilities.PaginationQuery) (*models.PricetypesList, error) {
	return s.pricetypesRepo.GetAllPricetypes(ctx, query)
}

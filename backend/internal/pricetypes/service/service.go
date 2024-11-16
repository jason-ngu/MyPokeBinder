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

func (s *pricetypesService) Create(ctx context.Context, newPricetype *models.PricetypeModel) (*models.PricetypeModel, error) {
	return s.pricetypesRepo.Create(ctx, newPricetype)
}

func (s *pricetypesService) GetByID(ctx context.Context, pricetypeID int) (*models.PricetypeModel, error) {
	return s.pricetypesRepo.GetByID(ctx, pricetypeID)
}

func (s *pricetypesService) Search(ctx context.Context, searchParams *models.PricetypeSearchParams, query *utilities.PaginationQuery) (*models.PricetypesList, error) {
	return s.pricetypesRepo.Search(ctx, searchParams, query)
}

func (s *pricetypesService) GetAllPricetypes(ctx context.Context, query *utilities.PaginationQuery) (*models.PricetypesList, error) {
	return s.pricetypesRepo.GetAllPricetypes(ctx, query)
}

package service

import (
	"backend/internal/models"
	"backend/internal/subtypes"
	"backend/pkg/utilities"
	"context"
)

type subtypesService struct {
	subtypesRepo subtypes.Repository
}

func NewSubtypesService(subtypesRepo subtypes.Repository) subtypes.Service {
	return &subtypesService{subtypesRepo: subtypesRepo}
}

func (s *subtypesService) Create(ctx context.Context, newSubtype *models.SubtypeEntity) (*models.SubtypeEntity, error) {
	return s.subtypesRepo.Create(ctx, newSubtype)
}

func (s *subtypesService) GetByID(ctx context.Context, subtypeID int) (*models.SubtypeModel, error) {
	return s.subtypesRepo.GetByID(ctx, subtypeID)
}

func (s *subtypesService) GetAllSubtypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SubtypesList, error) {
	return s.subtypesRepo.GetAllSubtypes(ctx, query)
}

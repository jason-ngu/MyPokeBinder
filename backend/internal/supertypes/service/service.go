package service

import (
	"backend/internal/models"
	"backend/internal/supertypes"
	"backend/pkg/utilities"
	"context"
)

type supertypesService struct {
	supertypesRepo supertypes.Repository
}

func NewSupertypesService(supertypesRepo supertypes.Repository) supertypes.Service {
	return &supertypesService{supertypesRepo: supertypesRepo}
}

func (s *supertypesService) Create(ctx context.Context, newSupertype *models.SupertypeEntity) (*models.SupertypeEntity, error) {
	return s.supertypesRepo.Create(ctx, newSupertype)
}

func (s *supertypesService) GetByID(ctx context.Context, supertypeID int) (*models.SupertypeModel, error) {
	return s.supertypesRepo.GetByID(ctx, supertypeID)
}

func (s *supertypesService) GetAllSupertypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SupertypesList, error) {
	return s.supertypesRepo.GetAllSupertypes(ctx, query)
}

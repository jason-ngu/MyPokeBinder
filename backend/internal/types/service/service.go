package service

import (
	"backend/internal/models"
	"backend/internal/types"
	"backend/pkg/utilities"
	"context"
)

type typesService struct {
	typesRepo types.Repository
}

func NewTypesService(typesRepo types.Repository) types.Service {
	return &typesService{typesRepo: typesRepo}
}

func (s *typesService) Create(ctx context.Context, newType *models.TypeModel) (*models.TypeModel, error) {
	return s.typesRepo.Create(ctx, newType)
}

func (s *typesService) GetByID(ctx context.Context, typeID int) (*models.TypeModel, error) {
	return s.typesRepo.GetByID(ctx, typeID)
}

func (s *typesService) Search(ctx context.Context, searchParams *models.TypeSearchParams, query *utilities.PaginationQuery) (*models.TypesList, error) {
	return s.typesRepo.Search(ctx, searchParams, query)
}

func (s *typesService) GetAllTypes(ctx context.Context, query *utilities.PaginationQuery) (*models.TypesList, error) {
	return s.typesRepo.GetAllTypes(ctx, query)
}

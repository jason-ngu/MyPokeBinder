package service

import (
	"backend/internal/models"
	"backend/internal/sets"
	"backend/pkg/utilities"
	"context"
)

type setsService struct {
	setsRepo sets.Repository
}

func NewSeriesService(setsRepo sets.Repository) sets.Service {
	return &setsService{setsRepo: setsRepo}
}

func (s *setsService) Create(ctx context.Context, newSeries *models.SetModel) (*models.SetModel, error) {
	return s.setsRepo.Create(ctx, newSeries)
}

func (s *setsService) GetByID(ctx context.Context, setsID int) (*models.SetModel, error) {
	return s.setsRepo.GetByID(ctx, setsID)
}

func (s *setsService) Search(ctx context.Context, searchParams *models.SetSearchParams, query *utilities.PaginationQuery) (*models.SetsList, error) {
	return s.setsRepo.Search(ctx, searchParams, query)
}

func (s *setsService) GetAllSets(ctx context.Context, query *utilities.PaginationQuery) (*models.SetsList, error) {
	return s.setsRepo.GetAllSets(ctx, query)
}

func (s *setsService) Delete(ctx context.Context, setID int) error {
	return s.setsRepo.Delete(ctx, setID)
}

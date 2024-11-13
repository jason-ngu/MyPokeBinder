package service

import (
	"backend/internal/models"
	"backend/internal/series"
	"backend/pkg/utilities"
	"context"
)

type seriesService struct {
	seriesRepo series.Repository
}

func NewSeriesService(seriesRepo series.Repository) series.Service {
	return &seriesService{seriesRepo: seriesRepo}
}

func (s *seriesService) Create(ctx context.Context, newSeries *models.SeriesModel) (*models.SeriesModel, error) {
	return s.seriesRepo.Create(ctx, newSeries)
}

func (s *seriesService) GetByID(ctx context.Context, seriesID int) (*models.SeriesModel, error) {
	return s.seriesRepo.GetByID(ctx, seriesID)
}

func (s *seriesService) SearchSeries(ctx context.Context, searchParams *models.SeriesSearchParams, query *utilities.PaginationQuery) (*models.SeriesList, error) {
	return s.seriesRepo.SearchSeries(ctx, searchParams, query)
}

func (s *seriesService) GetAllSeries(ctx context.Context, query *utilities.PaginationQuery) (*models.SeriesList, error) {
	return s.seriesRepo.GetAllSeries(ctx, query)
}

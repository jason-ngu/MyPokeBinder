package service

import (
	"backend/internal/models"
	"backend/internal/rarities"
	"backend/pkg/utilities"
	"context"
)

type raritiesService struct {
	raritiesRepo rarities.Repository
}

func NewRaritiesService(raritiesRepo rarities.Repository) rarities.Service {
	return &raritiesService{raritiesRepo: raritiesRepo}
}

func (s *raritiesService) Create(ctx context.Context, newRarity *models.RarityModel) (*models.RarityModel, error) {
	return s.raritiesRepo.Create(ctx, newRarity)
}

func (s *raritiesService) GetByID(ctx context.Context, rarityID int) (*models.RarityModel, error) {
	return s.raritiesRepo.GetByID(ctx, rarityID)
}

func (s *raritiesService) Search(ctx context.Context, searchParams *models.RaritySearchParams, query *utilities.PaginationQuery) (*models.RaritiesList, error) {
	return s.raritiesRepo.Search(ctx, searchParams, query)
}

func (s *raritiesService) GetAllRarities(ctx context.Context, query *utilities.PaginationQuery) (*models.RaritiesList, error) {
	return s.raritiesRepo.GetAllRarities(ctx, query)
}

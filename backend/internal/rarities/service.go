package rarities

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newRarity *models.RarityModel) (*models.RarityModel, error)
	GetByID(ctx context.Context, rarityID int) (*models.RarityModel, error)
	Search(ctx context.Context, searchParams *models.RaritySearchParams, query *utilities.PaginationQuery) (*models.RaritiesList, error)
	GetAllRarities(ctx context.Context, query *utilities.PaginationQuery) (*models.RaritiesList, error)
}

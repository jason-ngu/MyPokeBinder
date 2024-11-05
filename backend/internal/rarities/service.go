package rarities

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newRarity *models.RarityEntity) (*models.RarityEntity, error)
	GetByID(ctx context.Context, rarityID int) (*models.RarityModel, error)
	GetAllRarities(ctx context.Context, query *utilities.PaginationQuery) (*models.RaritiesList, error)
}

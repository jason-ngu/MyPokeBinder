package pricetypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newPricetype *models.PricetypeEntity) (*models.PricetypeEntity, error)
	GetByID(ctx context.Context, rarityID int) (*models.PricetypeModel, error)
	GetAllPricetypes(ctx context.Context, query *utilities.PaginationQuery) (*models.PricetypesList, error)
}

package pricetypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newPricetype *models.PricetypeModel) (*models.PricetypeModel, error)
	GetByID(ctx context.Context, rarityID int) (*models.PricetypeModel, error)
	Search(ctx context.Context, searchParams *models.PricetypeSearchParams, query *utilities.PaginationQuery) (*models.PricetypesList, error)
	GetAllPricetypes(ctx context.Context, query *utilities.PaginationQuery) (*models.PricetypesList, error)
}

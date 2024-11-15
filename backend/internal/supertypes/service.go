package supertypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newSupertype *models.SupertypeModel) (*models.SupertypeModel, error)
	GetByID(ctx context.Context, supertypeID int) (*models.SupertypeModel, error)
	Search(ctx context.Context, searchParams *models.SupertypeSearchParams, query *utilities.PaginationQuery) (*models.SupertypesList, error)
	GetAllSupertypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SupertypesList, error)
}

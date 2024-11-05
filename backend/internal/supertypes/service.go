package supertypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newSupertype *models.SupertypeEntity) (*models.SupertypeEntity, error)
	GetByID(ctx context.Context, supertypeID int) (*models.SupertypeModel, error)
	GetAllSupertypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SupertypesList, error)
}

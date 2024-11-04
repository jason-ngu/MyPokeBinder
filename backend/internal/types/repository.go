package types

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newType *models.TypeEntity) (*models.TypeEntity, error)
	GetByID(ctx context.Context, typeID int) (*models.TypeModel, error)
	GetAllTypes(ctx context.Context, query *utilities.PaginationQuery) (*models.TypesList, error)
}

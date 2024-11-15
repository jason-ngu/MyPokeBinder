package types

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newType *models.TypeModel) (*models.TypeModel, error)
	GetByID(ctx context.Context, typeID int) (*models.TypeModel, error)
	Search(ctx context.Context, searchParams *models.TypeSearchParams, query *utilities.PaginationQuery) (*models.TypesList, error)
	GetAllTypes(ctx context.Context, query *utilities.PaginationQuery) (*models.TypesList, error)
}

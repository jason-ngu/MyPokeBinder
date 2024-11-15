package subtypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newSubtype *models.SubtypeModel) (*models.SubtypeModel, error)
	GetByID(ctx context.Context, subtypeID int) (*models.SubtypeModel, error)
	Search(ctx context.Context, searchParams *models.SubtypeSearchParams, query *utilities.PaginationQuery) (*models.SubtypesList, error)
	GetAllSubtypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SubtypesList, error)
}

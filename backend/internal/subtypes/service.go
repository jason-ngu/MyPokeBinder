package subtypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newSubtype *models.SubtypeModel) (*models.SubtypeModel, error)
	GetByID(ctx context.Context, subtypeID int) (*models.SubtypeModel, error)
	SearchSubtypes(ctx context.Context, searchParams *models.SubtypeSearchParams, query *utilities.PaginationQuery) (*models.SubtypesList, error)
	GetAllSubtypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SubtypesList, error)
}

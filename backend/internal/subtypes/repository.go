package subtypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newSubtype *models.SubtypeEntity) (*models.SubtypeEntity, error)
	GetByID(ctx context.Context, subtypeID int) (*models.SubtypeModel, error)
	GetAllSubtypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SubtypesList, error)
}

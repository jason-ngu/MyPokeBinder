package sets

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newSet *models.SetModel) (*models.SetModel, error)
	GetByID(ctx context.Context, setID int) (*models.SetModel, error)
	Search(ctx context.Context, searchParams *models.SetSearchParams, query *utilities.PaginationQuery) (*models.SetsList, error)
	GetAllSets(ctx context.Context, query *utilities.PaginationQuery) (*models.SetsList, error)
	Delete(ctx context.Context, setID int) error
}

package sets

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newSet *models.SetEntity) (*models.SetEntity, error)
	GetByID(ctx context.Context, setID int) (*models.SetModel, error)
	GetAllSets(ctx context.Context, query *utilities.PaginationQuery) (*models.SetsList, error)
	Delete(ctx context.Context, setID int) error
}

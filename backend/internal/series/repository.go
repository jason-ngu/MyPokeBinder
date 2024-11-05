package series

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newSeries *models.SeriesEntity) (*models.SeriesEntity, error)
	GetByID(ctx context.Context, seriesID int) (*models.SeriesModel, error)
	GetAllSeries(ctx context.Context, query *utilities.PaginationQuery) (*models.SeriesList, error)
}

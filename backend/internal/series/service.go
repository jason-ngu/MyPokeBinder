package series

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newSeries *models.SeriesModel) (*models.SeriesModel, error)
	GetByID(ctx context.Context, seriesID int) (*models.SeriesModel, error)
	Search(ctx context.Context, searchParams *models.SeriesSearchParams, query *utilities.PaginationQuery) (*models.SeriesList, error)
	GetAllSeries(ctx context.Context, query *utilities.PaginationQuery) (*models.SeriesList, error)
}

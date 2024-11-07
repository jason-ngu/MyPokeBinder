package repository

import (
	"backend/internal/models"
	"backend/internal/series"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type seriesRepo struct {
	db *sql.DB
}

func NewSeriesRepository(db *sql.DB) series.Repository {
	return &seriesRepo{db: db}
}

func (r *seriesRepo) Create(ctx context.Context, newSeries *models.SeriesEntity) (*models.SeriesEntity, error) {
	series := &models.SeriesEntity{}
	row := r.db.QueryRowContext(ctx, createSeries, &newSeries.SeriesName)
	err := row.Scan(series)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.Create.QueryRowContext.Scan")
	}

	return series, nil
}

func (r *seriesRepo) GetByID(ctx context.Context, seriesID int) (*models.SeriesModel, error) {
	series := &models.SeriesModel{}
	err := r.db.QueryRowContext(ctx, getSeriesById, &seriesID).Scan(series)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetByID.QueryRowContext.Scan")
	}

	return series, nil
}

func (r *seriesRepo) GetAllSeries(ctx context.Context, query *utilities.PaginationQuery) (*models.SeriesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllSeries).Scan(totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.SeriesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.SeriesModel, 0),
		}, nil
	}

	var seriesList []*models.SeriesModel
	rows, err := r.db.QueryContext(ctx, getAllSeries)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.QueryContext")
	}
	for rows.Next() {
		var series models.SeriesModel
		err := rows.Scan(&series)
		if err != nil {
			return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.QueryContext.Scan")
		}
		seriesList = append(seriesList, &series)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.rows.Err")
	}

	return &models.SeriesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         seriesList,
	}, nil
}

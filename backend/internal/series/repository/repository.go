package repository

import (
	"backend/internal/models"
	"backend/internal/series"
	"backend/pkg/utilities"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type seriesRepo struct {
	db *sqlx.DB
}

func NewSeriesRepository(db *sqlx.DB) series.Repository {
	return &seriesRepo{db: db}
}

func (r *seriesRepo) Create(ctx context.Context, newSeries *models.SeriesModel) (*models.SeriesModel, error) {
	series := &models.SeriesEntity{}
	row := r.db.QueryRowxContext(ctx, createSeries, &newSeries.SeriesName)
	err := row.StructScan(series)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, series.SeriesID)
}

func (r *seriesRepo) GetByID(ctx context.Context, seriesID int) (*models.SeriesModel, error) {
	series := &models.SeriesModel{}
	err := r.db.GetContext(ctx, series, getSeriesById, seriesID)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetByID.GetContext")
	}

	return series, nil
}

func (r *seriesRepo) Search(ctx context.Context, searchParams *models.SeriesSearchParams, query *utilities.PaginationQuery) (*models.SeriesList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, getTotalCountAllSeries)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.Search.PrepareNamedContext.getTotalCountAllSeries")
	}
	err = nstmt.GetContext(ctx, &totalRecords, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.Search.GetContext")
	}
	if totalRecords == 0 {
		return &models.SeriesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.SeriesModel, 0),
		}, nil
	}

	var seriesList []models.SeriesModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(getAllSeries, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.Search.PrepareNamedContext.getAllSeries")
	}
	err = nstmt.SelectContext(ctx, &seriesList, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.Search.SelectContext")
	}

	return &models.SeriesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         seriesList,
	}, nil
}

func (r *seriesRepo) GetAllSeries(ctx context.Context, query *utilities.PaginationQuery) (*models.SeriesList, error) {
	var totalRecords int
	err := r.db.GetContext(ctx, &totalRecords, getTotalCountAllSeries)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.GetContext")
	}
	if totalRecords == 0 {
		return &models.SeriesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.SeriesModel, 0),
		}, nil
	}

	var seriesList []models.SeriesModel
	err = r.db.SelectContext(ctx, &seriesList, fmt.Sprintf(getAllSeries, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.SelectContext")
	}

	return &models.SeriesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         seriesList,
	}, nil
}

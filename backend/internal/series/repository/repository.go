package repository

import (
	"backend/internal/models"
	"backend/internal/series"
	"backend/pkg/utilities"
	"context"
	"fmt"
	"strings"

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
	row := r.db.QueryRowContext(ctx, createSeries, &newSeries.SeriesName)
	err := row.Scan(series)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.Create.QueryRowContext.Scan")
	}

	return &models.SeriesModel{
		SeriesID:   series.SeriesID,
		SeriesName: newSeries.SeriesName,
	}, nil
}

func (r *seriesRepo) GetByID(ctx context.Context, seriesID int) (*models.SeriesModel, error) {
	series := &models.SeriesModel{}
	err := r.db.QueryRowContext(ctx, getSeriesById, &seriesID).Scan(series)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetByID.QueryRowContext.Scan")
	}

	return series, nil
}

func (r *seriesRepo) SearchSeries(ctx context.Context, searchParams *models.SeriesSearchParams, query *utilities.PaginationQuery) (*models.SeriesList, error) {
	getAllSeriesCountWithSearchParams := getTotalCountAllSeries
	getAllSeriesWithSearchParams := getAllSeries

	var whereFilters []string
	if searchParams != nil {
		getAllSeriesCountWithSearchParams += " WHERE "
		getAllSeriesWithSearchParams += " WHERE "
		if searchParams.SeriesName != "" {
			formattedSeriesName := utilities.FormatStringForDatabase(searchParams.SeriesName)
			whereFilters = append(whereFilters, fmt.Sprintf("series_name = '%s'", formattedSeriesName))
		}
		getAllSeriesCountWithSearchParams += strings.Join(whereFilters, " AND ")
		getAllSeriesWithSearchParams += strings.Join(whereFilters, " AND ")
	}

	var totalRecords int
	err := r.db.QueryRowContext(ctx, getAllSeriesCountWithSearchParams).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.SearchSeries.QueryRowContext")
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
	rows, err := r.db.QueryContext(ctx, getAllSeriesWithSearchParams)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.SearchSeries.QueryContext")
	}
	for rows.Next() {
		var series models.SeriesModel
		err := rows.Scan(&series.SeriesID, &series.SeriesName)
		if err != nil {
			return nil, errors.Wrap(err, "seriesRepo.SearchSeries.QueryContext.Scan")
		}
		seriesList = append(seriesList, &series)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "seriesRepo.SearchSeries.rows.Err")
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
	err := r.db.QueryRowContext(ctx, getTotalCountAllSeries).Scan(&totalRecords)
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
		err := rows.Scan(&series.SeriesID, &series.SeriesName)
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

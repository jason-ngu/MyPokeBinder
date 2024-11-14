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
	row := r.db.QueryRowxContext(ctx, createSeries, &newSeries.SeriesName)
	err := row.StructScan(series)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.Create.QueryRowxContext.StructScan")
	}

	return &models.SeriesModel{
		SeriesID:   series.SeriesID,
		SeriesName: newSeries.SeriesName,
	}, nil
}

func (r *seriesRepo) GetByID(ctx context.Context, seriesID int) (*models.SeriesModel, error) {
	series := &models.SeriesModel{}
	row := r.db.QueryRowxContext(ctx, getSeriesById, seriesID)
	err := row.StructScan(series)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetByID.QueryRowxContext.StructScan")
	}

	return series, nil
}

func (r *seriesRepo) SearchSeries(ctx context.Context, searchParams *models.SeriesSearchParams, query *utilities.PaginationQuery) (*models.SeriesList, error) {
	var whereFilters []string
	var whereFiltersStr string
	if searchParams != nil {
		if searchParams.SeriesName != "" {
			formattedSeriesName := utilities.FormatStringForDatabase(searchParams.SeriesName)
			whereFilters = append(whereFilters, fmt.Sprintf("series_name = '%s'", formattedSeriesName))
		}
		whereFiltersStr = " WHERE " + strings.Join(whereFilters, " AND ")
	}
	getAllSeriesCountParams := fmt.Sprintf(getTotalCountAllSeries, whereFiltersStr)
	getAllSeriesParams := fmt.Sprintf(getAllSeries, whereFiltersStr, query.GetOffset(), query.GetLimit())

	var totalRecords int
	err := r.db.GetContext(ctx, &totalRecords, getAllSeriesCountParams)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.SearchSeries.GetContext")
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
	err = r.db.SelectContext(ctx, &seriesList, getAllSeriesParams)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.SearchSeries.SelectContext")
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
	err = r.db.SelectContext(ctx, &seriesList, getAllSeries, query.GetOffset(), query.GetLimit())
	// rows, err := r.db.QueryContext(ctx, getAllSeries)
	if err != nil {
		return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.SelectContext")
	}
	// for rows.Next() {
	// 	var series models.SeriesModel
	// 	err := rows.Scan(&series.SeriesID, &series.SeriesName)
	// 	if err != nil {
	// 		return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.QueryContext.Scan")
	// 	}
	// 	seriesList = append(seriesList, &series)
	// }
	// if err = rows.Err(); err != nil {
	// 	return nil, errors.Wrap(err, "seriesRepo.GetAllSeries.rows.Err")
	// }

	return &models.SeriesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         seriesList,
	}, nil
}

package repository

import (
	"backend/internal/models"
	"backend/internal/sets"
	"backend/pkg/utilities"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type setsRepo struct {
	db *sql.DB
}

func NewSetsRepository(db *sql.DB) sets.Repository {
	return &setsRepo{db: db}
}

func (r *setsRepo) Create(ctx context.Context, newSet *models.SetModel) (*models.SetModel, error) {
	syncDateCreated := time.Now()
	syncDateUpdated := time.Now()

	set := &models.SetEntity{}
	row := r.db.QueryRowContext(ctx, createSet, &newSet.SetCode, &newSet.SetName, &newSet.Series.SeriesID, &newSet.PtcgoCode, &newSet.CardTotal, &newSet.ExtendedCardTotal, &newSet.SetReleaseDate, &newSet.SymbolImage, &newSet.LogoImage, syncDateCreated, syncDateUpdated)
	err := row.Scan(set)
	if err != nil {
		return nil, errors.Wrap(err, "setRepo.Create.QueryRowContext.Scan")
	}

	return &models.SetModel{
		SetID:             set.SetID,
		SetCode:           newSet.SetCode,
		SetName:           newSet.SetName,
		Series:            newSet.Series,
		PtcgoCode:         newSet.PtcgoCode,
		CardTotal:         newSet.CardTotal,
		ExtendedCardTotal: newSet.ExtendedCardTotal,
		SetReleaseDate:    newSet.SetReleaseDate,
		SymbolImage:       newSet.SymbolImage,
		LogoImage:         newSet.LogoImage,
	}, nil
}

func (r *setsRepo) GetByID(ctx context.Context, setID int) (*models.SetModel, error) {
	set := &models.SetModel{}
	err := r.db.QueryRowContext(ctx, getSetById, setID).Scan(set)
	if err != nil {
		return nil, errors.Wrap(err, "setRepo.GetByID.QueryRowContext.Scan")
	}

	return set, nil
}

func (r *setsRepo) SearchSets(ctx context.Context, searchParams *models.SetSearchParams, query *utilities.PaginationQuery) (*models.SetsList, error) {
	getAllSetsCountWithSearchParams := getTotalCountAllSets
	getAllSetsWithSearchParams := getAllSets

	var whereFilters []string
	if searchParams != nil {
		getAllSetsCountWithSearchParams += " WHERE "
		getAllSetsWithSearchParams += " WHERE "
		if searchParams.SetName != "" {
			formattedSetName := utilities.FormatStringForDatabase(searchParams.SetName)
			whereFilters = append(whereFilters, fmt.Sprintf("set_name = '%s'", formattedSetName))
		}
		if searchParams.SetCode != "" {
			whereFilters = append(whereFilters, fmt.Sprintf("set_code = '%s'", searchParams.SetCode))
		}
		if searchParams.SeriesName != "" {
			whereFilters = append(whereFilters, fmt.Sprintf("series_name = '%s'", searchParams.SeriesName))
		}
		if searchParams.PtcgoCode != "" {
			whereFilters = append(whereFilters, fmt.Sprintf("ptcgo_code = '%s'", searchParams.PtcgoCode))
		}
		getAllSetsCountWithSearchParams += strings.Join(whereFilters, " AND ")
		getAllSetsWithSearchParams += strings.Join(whereFilters, " AND ")
	}

	var totalRecords int
	err := r.db.QueryRowContext(ctx, getAllSetsCountWithSearchParams).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.SearchSets.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.SetsList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.SetModel, 0),
		}, nil
	}

	var setsList []*models.SetModel
	rows, err := r.db.QueryContext(ctx, getAllSetsWithSearchParams)
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.SearchSets.QueryContext")
	}
	for rows.Next() {
		set := &models.SetModel{}
		err := rows.Scan(set)
		// err := rows.Scan(&set.SetID, &set.SetName, &set.SetCode, &set.Series, &set.PtcgoCode, &set.CardTotal, &set.ExtendedCardTotal, &set.SetReleaseDate, &set.SymbolImage, &set.LogoImage)
		if err != nil {
			return nil, errors.Wrap(err, "setsRepo.SearchSets.QueryContext.Scan")
		}
		setsList = append(setsList, set)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "setsRepo.SearchSets.rows.Err")
	}

	return &models.SetsList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         setsList,
	}, nil
}

func (r *setsRepo) GetAllSets(ctx context.Context, query *utilities.PaginationQuery) (*models.SetsList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllSets).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.GetAllSets.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.SetsList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.SetModel, 0),
		}, nil
	}

	var setsList []*models.SetModel
	rows, err := r.db.QueryContext(ctx, getAllSets)
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.GetAllSets.QueryContext")
	}
	for rows.Next() {
		var set models.SetModel
		err := rows.Scan(&set)
		if err != nil {
			return nil, errors.Wrap(err, "setsRepo.GetAllSets.QueryContext.Scan")
		}
		setsList = append(setsList, &set)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "setsRepo.GetAllSets.rows.Err")
	}

	return &models.SetsList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         setsList,
	}, nil
}

func (r *setsRepo) Delete(ctx context.Context, setID int) error {
	result, err := r.db.ExecContext(ctx, deleteSet, setID)
	if err != nil {
		return errors.Wrap(err, "setsRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "setsRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "setsRepo.Delete.rowsAffected")
	}

	return nil
}

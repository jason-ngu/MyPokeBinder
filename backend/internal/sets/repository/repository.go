package repository

import (
	"backend/internal/models"
	"backend/internal/sets"
	"backend/pkg/utilities"
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

type setsRepo struct {
	db *sql.DB
}

func NewSetsRepository(db *sql.DB) sets.Repository {
	return &setsRepo{db: db}
}

func (r *setsRepo) Create(ctx context.Context, newSet *models.SetEntity) (*models.SetEntity, error) {
	syncDateCreated := time.Now()
	syncDateUpdated := time.Now()

	set := &models.SetEntity{}
	row := r.db.QueryRowContext(ctx, createSet, &newSet.SetCode, &newSet.SetName, &newSet.SeriesID, &newSet.PtcgoCode, &newSet.CardTotal, &newSet.ExtendedCardTotal, &newSet.SetReleaseDate, &newSet.SymbolImage, &newSet.LogoImage, syncDateCreated, syncDateUpdated)
	err := row.Scan(set)
	if err != nil {
		return nil, errors.Wrap(err, "setRepo.Create.QueryRowContext.Scan")
	}

	return set, nil
}

func (r *setsRepo) GetByID(ctx context.Context, setID int) (*models.SetModel, error) {
	set := &models.SetModel{}
	err := r.db.QueryRowContext(ctx, getSetById, &setID).Scan(set)
	if err != nil {
		return nil, errors.Wrap(err, "setRepo.GetByID.QueryRowContext.Scan")
	}

	return set, nil
}

func (r *setsRepo) GetAllSets(ctx context.Context, query *utilities.PaginationQuery) (*models.SetsList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllSets).Scan(totalRecords)
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

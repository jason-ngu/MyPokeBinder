package repository

import (
	"backend/internal/models"
	"backend/internal/sets"
	"backend/pkg/utilities"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type setsRepo struct {
	db *sqlx.DB
}

func NewSetsRepository(db *sqlx.DB) sets.Repository {
	return &setsRepo{db: db}
}

func (r *setsRepo) Create(ctx context.Context, newSet *models.SetModel) (*models.SetModel, error) {
	syncDateCreated := time.Now()
	syncDateUpdated := time.Now()

	set := &models.SetEntity{}
	row := r.db.QueryRowxContext(ctx, createSet, &newSet.SetCode, &newSet.SetName, &newSet.Series.SeriesID, &newSet.PtcgoCode, &newSet.CardTotal, &newSet.ExtendedCardTotal, &newSet.SetReleaseDate, &newSet.SymbolImage, &newSet.LogoImage, syncDateCreated, syncDateUpdated)
	err := row.StructScan(set)
	if err != nil {
		return nil, errors.Wrap(err, "setRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, set.SetID)
}

func (r *setsRepo) GetByID(ctx context.Context, setID int) (*models.SetModel, error) {
	set := &models.SetModel{}
	err := r.db.GetContext(ctx, set, getSetById, setID)
	if err != nil {
		return nil, errors.Wrap(err, "setRepo.GetByID.GetContext")
	}

	return set, nil
}

func (r *setsRepo) Search(ctx context.Context, searchParams *models.SetSearchParams, query *utilities.PaginationQuery) (*models.SetsList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, getTotalCountAllSets)
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.Search.PrepareNamedContext.getTotalCountAllSets")
	}
	err = nstmt.GetContext(ctx, &totalRecords, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.Search.GetContext")
	}
	if totalRecords == 0 {
		return &models.SetsList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.SetModel, 0),
		}, nil
	}

	var setsList []models.SetModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(getAllSets, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.Search.PrepareNamedContext.getTotalCountAllSets")
	}
	err = nstmt.SelectContext(ctx, &setsList, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.Search.SelectContext")
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
	err := r.db.GetContext(ctx, &totalRecords, getTotalCountAllSets)
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.GetAllSets.GetContext")
	}
	if totalRecords == 0 {
		return &models.SetsList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.SetModel, 0),
		}, nil
	}

	var setsList []models.SetModel
	err = r.db.SelectContext(ctx, &setsList, fmt.Sprintf(getAllSets, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "setsRepo.GetAllSets.SelectContext")
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
		return errors.Wrap(sql.ErrNoRows, "setsRepo.Delete.rowsAffectedCount")
	}

	return nil
}

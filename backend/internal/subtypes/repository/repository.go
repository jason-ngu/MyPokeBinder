package repository

import (
	"backend/internal/models"
	"backend/internal/subtypes"
	"backend/pkg/utilities"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type subtypesRepo struct {
	db *sqlx.DB
}

func NewSubtypesRepository(db *sqlx.DB) subtypes.Repository {
	return &subtypesRepo{db: db}
}

func (r *subtypesRepo) Create(ctx context.Context, newSubtype *models.SubtypeModel) (*models.SubtypeModel, error) {
	subtype := &models.SubtypeEntity{}
	row := r.db.QueryRowxContext(ctx, createSubtype, &newSubtype.SubtypeName)
	err := row.StructScan(subtype)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, subtype.SubtypeID)
}

func (r *subtypesRepo) GetByID(ctx context.Context, subtypeID int) (*models.SubtypeModel, error) {
	subtype := &models.SubtypeModel{}
	err := r.db.GetContext(ctx, subtype, getSubtypeById, subtypeID)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetByID.GetContext")
	}

	return subtype, nil
}

func (r *subtypesRepo) Search(ctx context.Context, searchParams *models.SubtypeSearchParams, query *utilities.PaginationQuery) (*models.SubtypesList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, getTotalCountAllSubtypes)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.Search.PrepareNamedContext.getTotalCountAllSubtypes")
	}
	err = nstmt.GetContext(ctx, &totalRecords, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.Search.GetContext")
	}
	if totalRecords == 0 {
		return &models.SubtypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.SubtypeModel, 0),
		}, nil
	}

	var subtypesList []models.SubtypeModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(getAllSubtypes, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.Search.PrepareNamedContext.getAllSubtypes")
	}
	err = nstmt.SelectContext(ctx, &subtypesList, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.Search.SelectContext")
	}

	return &models.SubtypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         subtypesList,
	}, nil
}

func (r *subtypesRepo) GetAllSubtypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SubtypesList, error) {
	var totalRecords int
	err := r.db.GetContext(ctx, &totalRecords, getTotalCountAllSubtypes)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.GetContext")
	}
	if totalRecords == 0 {
		return &models.SubtypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.SubtypeModel, 0),
		}, nil
	}

	var subtypesList []models.SubtypeModel
	err = r.db.SelectContext(ctx, &subtypesList, fmt.Sprintf(getAllSubtypes, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.SelectContext")
	}

	return &models.SubtypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         subtypesList,
	}, nil
}

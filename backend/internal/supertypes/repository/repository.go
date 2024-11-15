package repository

import (
	"backend/internal/models"
	"backend/internal/supertypes"
	"backend/pkg/utilities"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type supertypesRepo struct {
	db *sqlx.DB
}

func NewSupertypesRepository(db *sqlx.DB) supertypes.Repository {
	return &supertypesRepo{db: db}
}

func (r *supertypesRepo) Create(ctx context.Context, newSupertype *models.SupertypeModel) (*models.SupertypeModel, error) {
	supertype := &models.SupertypeEntity{}
	row := r.db.QueryRowxContext(ctx, createSupertype, &newSupertype.SupertypeName)
	err := row.StructScan(supertype)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, supertype.SupertypeID)
}

func (r *supertypesRepo) GetByID(ctx context.Context, supertypeID int) (*models.SupertypeModel, error) {
	supertype := &models.SupertypeModel{}
	err := r.db.GetContext(ctx, supertype, getSupertypeById, supertypeID)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.GetByID.GetContext")
	}

	return supertype, nil
}

func (r *supertypesRepo) Search(ctx context.Context, searchParams *models.SupertypeSearchParams, query *utilities.PaginationQuery) (*models.SupertypesList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, getTotalCountAllSupertypes)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.Search.PrepareNamedContext.getTotalCountAllSupertypes")
	}
	err = nstmt.GetContext(ctx, &totalRecords, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.Search.GetContext")
	}
	if totalRecords == 0 {
		return &models.SupertypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.SupertypeModel, 0),
		}, nil
	}

	var supertypesList []models.SupertypeModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(getAllSupertypes, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.Search.PrepareNamedContext.getAllSupertypes")
	}
	err = nstmt.SelectContext(ctx, &supertypesList, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.Search.SelectContext")
	}

	return &models.SupertypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         supertypesList,
	}, nil
}

func (r *supertypesRepo) GetAllSupertypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SupertypesList, error) {
	var totalRecords int
	err := r.db.GetContext(ctx, &totalRecords, getTotalCountAllSupertypes)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.GetAllSupertypes.GetContext")
	}
	if totalRecords == 0 {
		return &models.SupertypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.SupertypeModel, 0),
		}, nil
	}

	var supertypesList []models.SupertypeModel
	err = r.db.SelectContext(ctx, &supertypesList, fmt.Sprintf(getAllSupertypes, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.GetAllSupertypes.SelectContext")
	}

	return &models.SupertypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         supertypesList,
	}, nil
}

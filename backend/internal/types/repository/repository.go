package repository

import (
	"backend/internal/models"
	"backend/internal/types"
	"backend/pkg/utilities"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type typesRepo struct {
	db *sqlx.DB
}

func NewTypesRepository(db *sqlx.DB) types.Repository {
	return &typesRepo{db: db}
}

func (r *typesRepo) Create(ctx context.Context, newType *models.TypeModel) (*models.TypeModel, error) {
	t := &models.TypeEntity{}
	row := r.db.QueryRowxContext(ctx, createType, &newType.TypeName)
	err := row.StructScan(t)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, t.TypeID)
}

func (r *typesRepo) GetByID(ctx context.Context, typeID int) (*models.TypeModel, error) {
	t := &models.TypeModel{}
	err := r.db.GetContext(ctx, t, getTypeById, typeID)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.GetByID.GetContext")
	}

	return t, nil
}

func (r *typesRepo) Search(ctx context.Context, searchParams *models.TypeSearchParams, query *utilities.PaginationQuery) (*models.TypesList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, utilities.FormatSqlQueryWithSearchParams(getTotalCountAllTypes, *searchParams, false))
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.Search.PrepareNamedContext.getTotalCountAllTypes")
	}
	err = nstmt.GetContext(ctx, &totalRecords, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.Search.GetContext")
	}
	if totalRecords == 0 {
		return &models.TypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.TypeModel, 0),
		}, nil
	}

	var typesList []models.TypeModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(utilities.FormatSqlQueryWithSearchParams(getAllTypes, *searchParams, true), query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.Search.PrepareNamedContext.getAllTypes")
	}
	err = nstmt.SelectContext(ctx, &typesList, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.Search.SelectContext")
	}

	return &models.TypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         typesList,
	}, nil
}

func (r *typesRepo) GetAllTypes(ctx context.Context, query *utilities.PaginationQuery) (*models.TypesList, error) {
	var totalRecords int
	err := r.db.GetContext(ctx, &totalRecords, getTotalCountAllTypes)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.GetAllTypes.GetContext")
	}
	if totalRecords == 0 {
		return &models.TypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.TypeModel, 0),
		}, nil
	}

	var typesList []models.TypeModel
	err = r.db.SelectContext(ctx, &typesList, fmt.Sprintf(getAllTypes, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.GetAllTypes.SelectContext")
	}

	return &models.TypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         typesList,
	}, nil
}

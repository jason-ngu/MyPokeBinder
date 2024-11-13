package repository

import (
	"backend/internal/models"
	"backend/internal/types"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type typesRepo struct {
	db *sql.DB
}

func NewTypesRepository(db *sql.DB) types.Repository {
	return &typesRepo{db: db}
}

func (r *typesRepo) Create(ctx context.Context, newType *models.TypeEntity) (*models.TypeEntity, error) {
	t := &models.TypeEntity{}
	row := r.db.QueryRowContext(ctx, createType, &newType.TypeName)
	err := row.Scan(t)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.Create.QueryRowContext.Scan")
	}

	return t, nil
}

func (r *typesRepo) GetByID(ctx context.Context, typeID int) (*models.TypeModel, error) {
	t := &models.TypeModel{}
	err := r.db.QueryRowContext(ctx, getTypeById, &typeID).Scan(t)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.GetByID.QueryRowContext.Scan")
	}

	return t, nil
}

func (r *typesRepo) GetAllTypes(ctx context.Context, query *utilities.PaginationQuery) (*models.TypesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllTypes).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.GetAllTypes.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.TypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.TypeModel, 0),
		}, nil
	}

	var typesList []*models.TypeModel
	rows, err := r.db.QueryContext(ctx, getAllTypes)
	if err != nil {
		return nil, errors.Wrap(err, "typesRepo.GetAllTypes.QueryContext")
	}
	for rows.Next() {
		var t models.TypeModel
		err := rows.Scan(&t)
		if err != nil {
			return nil, errors.Wrap(err, "typesRepo.GetAllTypes.QueryContext.Scan")
		}
		typesList = append(typesList, &t)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "typesRepo.GetAllTypes.rows.Err")
	}

	return &models.TypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         typesList,
	}, nil
}

package repository

import (
	"backend/internal/models"
	"backend/internal/supertypes"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type supertypesRepo struct {
	db *sql.DB
}

func NewSupertypesRepository(db *sql.DB) supertypes.Repository {
	return &supertypesRepo{db: db}
}

func (r *supertypesRepo) Create(ctx context.Context, newSupertype *models.SupertypeEntity) (*models.SupertypeEntity, error) {
	supertype := &models.SupertypeEntity{}
	row := r.db.QueryRowContext(ctx, createSupertype, &newSupertype.SupertypeName)
	err := row.Scan(supertype)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.Create.QueryRowContext.Scan")
	}

	return supertype, nil
}

func (r *supertypesRepo) GetByID(ctx context.Context, supertypeID int) (*models.SupertypeModel, error) {
	supertype := &models.SupertypeModel{}
	err := r.db.QueryRowContext(ctx, getSupertypeById, &supertypeID).Scan(supertype)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.GetByID.QueryRowContext.Scan")
	}

	return supertype, nil
}

func (r *supertypesRepo) GetAllSupertypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SupertypesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllSupertypes).Scan(totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.GetAllSupertypes.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.SupertypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.SupertypeModel, 0),
		}, nil
	}

	var supertypesList []*models.SupertypeModel
	rows, err := r.db.QueryContext(ctx, getAllSupertypes)
	if err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.GetAllSupertypes.QueryContext")
	}
	for rows.Next() {
		var supertype models.SupertypeModel
		err := rows.Scan(&supertype)
		if err != nil {
			return nil, errors.Wrap(err, "supertypesRepo.GetAllSupertypes.QueryContext.Scan")
		}
		supertypesList = append(supertypesList, &supertype)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "supertypesRepo.GetAllSupertypes.rows.Err")
	}

	return &models.SupertypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         supertypesList,
	}, nil
}

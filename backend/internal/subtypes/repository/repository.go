package repository

import (
	"backend/internal/models"
	"backend/internal/subtypes"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type subtypesRepo struct {
	db *sql.DB
}

func NewSubtypesRepository(db *sql.DB) subtypes.Repository {
	return &subtypesRepo{db: db}
}

func (r *subtypesRepo) Create(ctx context.Context, newSubtype *models.SubtypeEntity) (*models.SubtypeEntity, error) {
	subtype := &models.SubtypeEntity{}
	row := r.db.QueryRowContext(ctx, createSubtype, &newSubtype.SubtypeName)
	err := row.Scan(subtype)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.Create.QueryRowContext.Scan")
	}

	return subtype, nil
}

func (r *subtypesRepo) GetByID(ctx context.Context, subtypeID int) (*models.SubtypeModel, error) {
	subtype := &models.SubtypeModel{}
	err := r.db.QueryRowContext(ctx, getSubtypeById, &subtypeID).Scan(subtype)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetByID.QueryRowContext.Scan")
	}

	return subtype, nil
}

func (r *subtypesRepo) GetAllSubtypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SubtypesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllSubtypes).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.SubtypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.SubtypeModel, 0),
		}, nil
	}

	var subtypesList []*models.SubtypeModel
	rows, err := r.db.QueryContext(ctx, getAllSubtypes)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.QueryContext")
	}
	for rows.Next() {
		var subtype models.SubtypeModel
		err := rows.Scan(&subtype)
		if err != nil {
			return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.QueryContext.Scan")
		}
		subtypesList = append(subtypesList, &subtype)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.rows.Err")
	}

	return &models.SubtypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         subtypesList,
	}, nil
}

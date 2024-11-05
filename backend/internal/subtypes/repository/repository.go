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
	t := &models.SubtypeEntity{}
	row := r.db.QueryRowContext(ctx, createSubtype, &newSubtype.SubtypeName)
	err := row.Scan(t)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.Create.QueryRowContext.Scan")
	}

	return t, nil
}

func (r *subtypesRepo) GetByID(ctx context.Context, subtypeID int) (*models.SubtypeModel, error) {
	t := &models.SubtypeModel{}
	err := r.db.QueryRowContext(ctx, getSubtypeById, &subtypeID).Scan(t)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetByID.QueryRowContext.Scan")
	}

	return t, nil
}

func (r *subtypesRepo) GetAllSubtypes(ctx context.Context, query *utilities.PaginationQuery) (*models.SubtypesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllSubtypes).Scan(totalRecords)
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

	var typesList []*models.SubtypeModel
	rows, err := r.db.QueryContext(ctx, getAllSubtypes)
	if err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.QueryContext")
	}
	for rows.Next() {
		var t models.SubtypeModel
		err := rows.Scan(&t)
		if err != nil {
			return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.QueryContext.Scan")
		}
		typesList = append(typesList, &t)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "subtypesRepo.GetAllSubtypes.rows.Err")
	}

	return &models.SubtypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         typesList,
	}, nil
}

package repository

import (
	"backend/internal/models"
	"backend/internal/pricetypes"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type pricetypesRepo struct {
	db *sql.DB
}

func NewPricetypesRepository(db *sql.DB) pricetypes.Repository {
	return &pricetypesRepo{db: db}
}

func (r *pricetypesRepo) Create(ctx context.Context, newPricetype *models.PricetypeEntity) (*models.PricetypeEntity, error) {
	t := &models.PricetypeEntity{}
	row := r.db.QueryRowContext(ctx, createPricetype, &newPricetype.PricetypeName)
	err := row.Scan(t)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.Create.QueryRowContext.Scan")
	}

	return t, nil
}

func (r *pricetypesRepo) GetByID(ctx context.Context, pricetypeID int) (*models.PricetypeModel, error) {
	t := &models.PricetypeModel{}
	err := r.db.QueryRowContext(ctx, getPricetypeById, &pricetypeID).Scan(t)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.GetByID.QueryRowContext.Scan")
	}

	return t, nil
}

func (r *pricetypesRepo) GetAllPricetypes(ctx context.Context, query *utilities.PaginationQuery) (*models.PricetypesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllPricetypes).Scan(totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.GetAllPricetypes.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.PricetypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.PricetypeModel, 0),
		}, nil
	}

	var typesList []*models.PricetypeModel
	rows, err := r.db.QueryContext(ctx, getAllPricetypes)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.GetAllPricetypes.QueryContext")
	}
	for rows.Next() {
		var t models.PricetypeModel
		err := rows.Scan(&t)
		if err != nil {
			return nil, errors.Wrap(err, "pricetypesRepo.GetAllPricetypes.QueryContext.Scan")
		}
		typesList = append(typesList, &t)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.GetAllPricetypes.rows.Err")
	}

	return &models.PricetypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         typesList,
	}, nil
}

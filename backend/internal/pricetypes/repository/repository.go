package repository

import (
	"backend/internal/models"
	"backend/internal/pricetypes"
	"backend/pkg/utilities"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type pricetypesRepo struct {
	db *sqlx.DB
}

func NewPricetypesRepository(db *sqlx.DB) pricetypes.Repository {
	return &pricetypesRepo{db: db}
}

func (r *pricetypesRepo) Create(ctx context.Context, newPricetype *models.PricetypeModel) (*models.PricetypeModel, error) {
	pricetype := &models.PricetypeEntity{}
	row := r.db.QueryRowxContext(ctx, createPricetype, &newPricetype.PricetypeName)
	err := row.StructScan(pricetype)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, pricetype.PricetypeID)
}

func (r *pricetypesRepo) GetByID(ctx context.Context, pricetypeID int) (*models.PricetypeModel, error) {
	pricetype := &models.PricetypeModel{}
	err := r.db.GetContext(ctx, pricetype, getPricetypeById, pricetypeID)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.GetByID.GetContext")
	}

	return pricetype, nil
}

func (r *pricetypesRepo) Search(ctx context.Context, searchParams *models.PricetypeSearchParams, query *utilities.PaginationQuery) (*models.PricetypesList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, utilities.FormatSqlQueryWithSearchParams(getTotalCountAllPricetypes, *searchParams, false))
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.Search.PrepareNamedContext")
	}
	err = nstmt.GetContext(ctx, &totalRecords, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.Search.GetContext")
	}
	if totalRecords == 0 {
		return &models.PricetypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.PricetypeModel, 0),
		}, nil
	}

	var pricetypesList []models.PricetypeModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(utilities.FormatSqlQueryWithSearchParams(getAllPricetypes, *searchParams, true), query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.Search.PrepareNamedContext.getAllPricetypes")
	}
	err = nstmt.SelectContext(ctx, &pricetypesList, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.Search.SelectContext")
	}

	return &models.PricetypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         pricetypesList,
	}, nil
}

func (r *pricetypesRepo) GetAllPricetypes(ctx context.Context, query *utilities.PaginationQuery) (*models.PricetypesList, error) {
	var totalRecords int
	err := r.db.GetContext(ctx, &totalRecords, getTotalCountAllPricetypes)
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.GetAllPricetypes.GetContext")
	}
	if totalRecords == 0 {
		return &models.PricetypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.PricetypeModel, 0),
		}, nil
	}

	var pricetypesList []models.PricetypeModel
	err = r.db.SelectContext(ctx, &pricetypesList, fmt.Sprintf(getAllPricetypes, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "pricetypesRepo.GetAllPricetypes.SelectContext")
	}

	return &models.PricetypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         pricetypesList,
	}, nil
}

package repository

import (
	"backend/internal/models"
	"backend/internal/rarities"
	"backend/pkg/utilities"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type raritiesRepo struct {
	db *sqlx.DB
}

func NewRaritiesRepository(db *sqlx.DB) rarities.Repository {
	return &raritiesRepo{db: db}
}

func (r *raritiesRepo) Create(ctx context.Context, newRarity *models.RarityModel) (*models.RarityModel, error) {
	rarity := &models.RarityEntity{}
	row := r.db.QueryRowxContext(ctx, createRarity, &newRarity.RarityName)
	err := row.StructScan(rarity)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, rarity.RarityID)
}

func (r *raritiesRepo) GetByID(ctx context.Context, rarityID int) (*models.RarityModel, error) {
	rarity := &models.RarityModel{}
	err := r.db.GetContext(ctx, rarity, getRarityById, rarityID)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetByID.GetContext")
	}

	return rarity, nil
}

func (r *raritiesRepo) Search(ctx context.Context, searchParams *models.RaritySearchParams, query *utilities.PaginationQuery) (*models.RaritiesList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, getTotalCountAllRarities)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.Search.PrepareNamedContext.getTotalCountAllRarities")
	}
	err = nstmt.GetContext(ctx, &totalRecords, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.Search.GetContext")
	}
	if totalRecords == 0 {
		return &models.RaritiesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.RarityModel, 0),
		}, nil
	}

	var raritiesList []models.RarityModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(getAllRarities, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.Search.PrepareNamedContext.getAllRarities")
	}
	err = nstmt.SelectContext(ctx, &raritiesList, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.Search.SelectContext")
	}

	return &models.RaritiesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         raritiesList,
	}, nil
}

func (r *raritiesRepo) GetAllRarities(ctx context.Context, query *utilities.PaginationQuery) (*models.RaritiesList, error) {
	var totalRecords int
	err := r.db.GetContext(ctx, &totalRecords, getTotalCountAllRarities)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.GetContext")
	}
	if totalRecords == 0 {
		return &models.RaritiesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.RarityModel, 0),
		}, nil
	}

	var raritiesList []models.RarityModel
	err = r.db.SelectContext(ctx, &raritiesList, fmt.Sprintf(getAllRarities, query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.SelectContext")
	}

	return &models.RaritiesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         raritiesList,
	}, nil
}

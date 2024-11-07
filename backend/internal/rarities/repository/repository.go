package repository

import (
	"backend/internal/models"
	"backend/internal/rarities"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type raritiesRepo struct {
	db *sql.DB
}

func NewRaritiesRepository(db *sql.DB) rarities.Repository {
	return &raritiesRepo{db: db}
}

func (r *raritiesRepo) Create(ctx context.Context, newRarity *models.RarityEntity) (*models.RarityEntity, error) {
	rarity := &models.RarityEntity{}
	row := r.db.QueryRowContext(ctx, createRarity, &newRarity.RarityName)
	err := row.Scan(rarity)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.Create.QueryRowContext.Scan")
	}

	return rarity, nil
}

func (r *raritiesRepo) GetByID(ctx context.Context, rarityID int) (*models.RarityModel, error) {
	rarity := &models.RarityModel{}
	err := r.db.QueryRowContext(ctx, getRarityById, &rarityID).Scan(rarity)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetByID.QueryRowContext.Scan")
	}

	return rarity, nil
}

func (r *raritiesRepo) GetAllRarities(ctx context.Context, query *utilities.PaginationQuery) (*models.RaritiesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountAllRarities).Scan(totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.RaritiesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.RarityModel, 0),
		}, nil
	}

	var raritiesList []*models.RarityModel
	rows, err := r.db.QueryContext(ctx, getAllRarities)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.QueryContext")
	}
	for rows.Next() {
		var rarity models.RarityModel
		err := rows.Scan(&rarity)
		if err != nil {
			return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.QueryContext.Scan")
		}
		raritiesList = append(raritiesList, &rarity)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.rows.Err")
	}

	return &models.RaritiesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         raritiesList,
	}, nil
}

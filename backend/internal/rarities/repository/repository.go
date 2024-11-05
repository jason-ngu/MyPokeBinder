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
	t := &models.RarityEntity{}
	row := r.db.QueryRowContext(ctx, createRarity, &newRarity.RarityName)
	err := row.Scan(t)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.Create.QueryRowContext.Scan")
	}

	return t, nil
}

func (r *raritiesRepo) GetByID(ctx context.Context, subtypeID int) (*models.RarityModel, error) {
	t := &models.RarityModel{}
	err := r.db.QueryRowContext(ctx, getRarityById, &subtypeID).Scan(t)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetByID.QueryRowContext.Scan")
	}

	return t, nil
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

	var typesList []*models.RarityModel
	rows, err := r.db.QueryContext(ctx, getAllRarities)
	if err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.QueryContext")
	}
	for rows.Next() {
		var t models.RarityModel
		err := rows.Scan(&t)
		if err != nil {
			return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.QueryContext.Scan")
		}
		typesList = append(typesList, &t)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "raritiesRepo.GetAllRarities.rows.Err")
	}

	return &models.RaritiesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         typesList,
	}, nil
}

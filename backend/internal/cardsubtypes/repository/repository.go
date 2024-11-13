package repository

import (
	"backend/internal/cardsubtypes"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type cardsubtypesRepo struct {
	db *sql.DB
}

func NewCardSubtypesRepository(db *sql.DB) cardsubtypes.Repository {
	return &cardsubtypesRepo{db: db}
}

func (r *cardsubtypesRepo) Create(ctx context.Context, newCardSubtype *models.CardSubtypeEntity) (*models.CardSubtypeEntity, error) {
	cardSubtype := &models.CardSubtypeEntity{}
	row := r.db.QueryRowContext(ctx, createCardSubtype, &newCardSubtype.CardID, &newCardSubtype.SubtypeID)
	err := row.Scan(cardSubtype)
	if err != nil {
		return nil, errors.Wrap(err, "cardsubtypesRepo.Create.QueryRowContext.Scan")
	}

	return cardSubtype, nil
}

func (r *cardsubtypesRepo) GetByCardID(ctx context.Context, cardID int, query *utilities.PaginationQuery) (*models.CardSubtypesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountByCardID, cardID).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "cardsubtypesRepo.GetByCardID.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.CardSubtypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.CardSubtypeModel, 0),
		}, nil
	}

	var cardsubtypesList []*models.CardSubtypeModel
	rows, err := r.db.QueryContext(ctx, getCardSubtypeByCardID, cardID)
	if err != nil {
		return nil, errors.Wrap(err, "cardsubtypesRepo.GetByCardID.QueryContext")
	}
	for rows.Next() {
		var cardsubtype models.CardSubtypeModel
		err := rows.Scan(&cardsubtype)
		if err != nil {
			return nil, errors.Wrap(err, "cardsubtypesRepo.GetByCardID.QueryContext.Scan")
		}
		cardsubtypesList = append(cardsubtypesList, &cardsubtype)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "cardsubtypesRepo.GetByCardID.rows.Err")
	}

	return &models.CardSubtypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         cardsubtypesList,
	}, nil
}

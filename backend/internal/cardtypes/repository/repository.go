package repository

import (
	"backend/internal/cardtypes"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type cardtypesRepo struct {
	db *sql.DB
}

func NewCardtypesRepository(db *sql.DB) cardtypes.Repository {
	return &cardtypesRepo{db: db}
}

func (r *cardtypesRepo) Create(ctx context.Context, newCardtype *models.CardTypeEntity) (*models.CardTypeEntity, error) {
	cardtype := &models.CardTypeEntity{}
	row := r.db.QueryRowContext(ctx, createCardType, &newCardtype.CardID, &newCardtype.TypeID)
	err := row.Scan(cardtype)
	if err != nil {
		return nil, errors.Wrap(err, "cardtypesRepo.Create.QueryRowContext.Scan")
	}

	return cardtype, nil
}

func (r *cardtypesRepo) GetByCardID(ctx context.Context, cardID int, query *utilities.PaginationQuery) (*models.CardTypesList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountByCardID, cardID).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "cardtypesRepo.GetByCardID.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.CardTypesList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.CardTypeModel, 0),
		}, nil
	}

	var cardtypesList []*models.CardTypeModel
	rows, err := r.db.QueryContext(ctx, getCardTypeByCardID, cardID)
	if err != nil {
		return nil, errors.Wrap(err, "cardtypesRepo.GetByCardID.QueryContext")
	}
	for rows.Next() {
		var cardtype models.CardTypeModel
		err := rows.Scan(&cardtype)
		if err != nil {
			return nil, errors.Wrap(err, "cardtypesRepo.GetByCardID.QueryContext.Scan")
		}
		cardtypesList = append(cardtypesList, &cardtype)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "cardtypesRepo.GetByCardID.rows.Err")
	}

	return &models.CardTypesList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         cardtypesList,
	}, nil
}

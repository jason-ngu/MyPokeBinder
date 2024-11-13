package repository

import (
	"backend/internal/collectioncards"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type collectioncardsRepo struct {
	db *sqlx.DB
}

func NewCollectionCardsRepository(db *sqlx.DB) collectioncards.Repository {
	return &collectioncardsRepo{db: db}
}

func (r *collectioncardsRepo) AddCardsToCollection(ctx context.Context, collectionId int, collectioncardsToAdd []*models.CollectionCardsModel) (*models.CollectionCardsList, error) {
	var strValues []string
	for _, collectioncard := range collectioncardsToAdd {
		strValues = append(strValues, fmt.Sprintf("(%d, %d, %d, %d, %s)",
			collectionId, &collectioncard.Card.CardID, &collectioncard.Quantity, &collectioncard.Grade, &collectioncard.GradingCompany))

		_, err := r.db.Exec(addCardsToCollection, strValues, &collectioncard.Quantity)
		if err != nil {
			return nil, err
		}
	}

	return r.GetByCollectionID(ctx, collectionId, &utilities.PaginationQuery{})
}

func (r *collectioncardsRepo) GetByCollectionID(ctx context.Context, collectionId int, query *utilities.PaginationQuery) (*models.CollectionCardsList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountByCollectionID, collectionId).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "collectioncardsRepo.GetByCollectionID.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.CollectionCardsList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.CollectionCardsModel, 0),
		}, nil
	}

	var collectioncardsList []*models.CollectionCardsModel
	rows, err := r.db.QueryContext(ctx, getByCollectionID, collectionId)
	if err != nil {
		return nil, errors.Wrap(err, "collectioncardsRepo.GetByCollectionID.QueryContext")
	}
	for rows.Next() {
		var collectioncard models.CollectionCardsModel
		err := rows.Scan(&collectioncard)
		if err != nil {
			return nil, errors.Wrap(err, "collectioncardsRepo.GetByCollectionID.QueryContext.Scan")
		}
		collectioncardsList = append(collectioncardsList, &collectioncard)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "collectioncardsRepo.GetByCollectionID.rows.Err")
	}

	return &models.CollectionCardsList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         collectioncardsList,
	}, nil
}

func (r *collectioncardsRepo) RemoveCardsFromCollection(ctx context.Context, collectionId int, collectioncardsToRemove []*models.CollectionCardsModel) (*models.CollectionCardsList, error) {
	for _, collectioncard := range collectioncardsToRemove {
		_, err := r.db.Exec(removeCardsFromCollection,
			&collectioncard.Quantity, &collectioncard.Card.CardID, &collectioncard.Grade, &collectioncard.GradingCompany)
		if err != nil {
			return nil, err
		}
	}

	return r.GetByCollectionID(ctx, collectionId, &utilities.PaginationQuery{})
}

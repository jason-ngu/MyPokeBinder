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

func (r *collectioncardsRepo) AddCardsToCollection(ctx context.Context, collectionID int, collectioncardsToAdd []*models.CollectionCardsModel) (*models.CollectionCardsList, error) {
	var strValues []string
	for _, collectioncard := range collectioncardsToAdd {
		strValues = append(strValues, fmt.Sprintf("(%d, %d, %d, %d, %s)",
			collectionID, collectioncard.Card.CardID, collectioncard.Quantity, collectioncard.Grade, collectioncard.GradingCompany))

		_, err := r.db.ExecContext(ctx, addCardsToCollection, strValues, &collectioncard.Quantity)
		if err != nil {
			return nil, errors.Wrap(err, "collectioncardsRepo.AddCardsToCollection.ExecContext")
		}
	}

	return r.GetByCollectionID(ctx, collectionID, &utilities.PaginationQuery{})
}

func (r *collectioncardsRepo) GetByCollectionID(ctx context.Context, collectionID int, query *utilities.PaginationQuery) (*models.CollectionCardsList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, utilities.FormatSqlQueryWithSearchParams(getTotalCountByCollectionID, nil, false))
	if err != nil {
		return nil, errors.Wrap(err, "collectioncardsRepo.GetByCollectionID.PrepareNamedContext.getTotalCountByCollectionID")
	}
	err = nstmt.GetContext(ctx, &totalRecords, collectionID)
	if err != nil {
		return nil, errors.Wrap(err, "collectioncardsRepo.GetByCollectionID.GetContext")
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
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(utilities.FormatSqlQueryWithSearchParams(getByCollectionID, nil, true), query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "collectioncardsRepo.Search.PrepareNamedContext.getByCollectionID")
	}
	err = nstmt.SelectContext(ctx, &collectioncardsList, collectionID)
	if err != nil {
		return nil, errors.Wrap(err, "collectioncardsRepo.Search.SelectContext")
	}

	return &models.CollectionCardsList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         collectioncardsList,
	}, nil
}

func (r *collectioncardsRepo) RemoveCardsFromCollection(ctx context.Context, collectionID int, collectioncardsToRemove []*models.CollectionCardsModel) (*models.CollectionCardsList, error) {
	for _, collectioncard := range collectioncardsToRemove {
		_, err := r.db.ExecContext(ctx, removeCardsFromCollection,
			collectioncard.Quantity, collectioncard.Card.CardID, collectioncard.Grade, collectioncard.GradingCompany)
		if err != nil {
			return nil, errors.Wrap(err, "collectioncardsRepo.RemoveCardsFromCollection.ExecContext")
		}
	}

	return r.GetByCollectionID(ctx, collectionID, &utilities.PaginationQuery{})
}

package repository

import (
	"backend/internal/collections"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type collectionsRepo struct {
	db *sqlx.DB
}

func NewCardsRepository(db *sqlx.DB) collections.Repository {
	return &collectionsRepo{db: db}
}

func (r *collectionsRepo) Create(ctx context.Context, newCollections *models.CollectionModel) (*models.CollectionModel, error) {
	collection := &models.CollectionEntity{}
	row := r.db.QueryRowxContext(ctx, createCollection, &newCollections.CollectionName, &newCollections.UserID)
	err := row.StructScan(collection)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, collection.CollectionID)
}

func (r *collectionsRepo) GetByID(ctx context.Context, collectionID int) (*models.CollectionModel, error) {
	collection := &models.CollectionModel{}
	err := r.db.GetContext(ctx, collection, getCollectionById, collectionID)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetByID.GetContext")
	}

	return collection, nil
}

func (r *collectionsRepo) GetAllCollectionsByUserID(ctx context.Context, userID int, query *utilities.PaginationQuery) (*models.CollectionsList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, utilities.FormatSqlQueryWithSearchParams(getTotalCountByUserID, nil, false))
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetAllCollectionsByUserID.PrepareNamedContext.getTotalCountByUserID")
	}
	err = nstmt.GetContext(ctx, &totalRecords, userID)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetAllCollectionsByUserID.GetContext")
	}
	if totalRecords == 0 {
		return &models.CollectionsList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.CollectionModel, 0),
		}, nil
	}

	var collectionsList []*models.CollectionModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(utilities.FormatSqlQueryWithSearchParams(getCollectionsByUserID, nil, true), query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetAllCollectionsByUserID.PrepareNamedContext.getCollectionsByUserID")
	}
	err = nstmt.SelectContext(ctx, &collectionsList, userID)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetAllCollectionsByUserID.SelectContext")
	}

	return &models.CollectionsList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         collectionsList,
	}, nil
}

func (r *collectionsRepo) Update(ctx context.Context, collectionToUpdate *models.CollectionModel) (*models.CollectionModel, error) {
	collection := &models.CollectionEntity{}
	row := r.db.QueryRowxContext(ctx, updateCollection, &collectionToUpdate.CollectionName, &collectionToUpdate.CollectionID, &collectionToUpdate.UserID)
	err := row.StructScan(collection)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.Update.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, collection.CollectionID)
}

func (r *collectionsRepo) Delete(ctx context.Context, collectionId int) error {
	result, err := r.db.ExecContext(ctx, deleteCollection, collectionId)
	if err != nil {
		return errors.Wrap(err, "collectionsRepo.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "collectionsRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "collectionsRepo.Delete.rowsAffected")
	}

	return nil
}

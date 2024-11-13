package repository

import (
	"backend/internal/collections"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type collectionsRepo struct {
	db *sql.DB
}

func NewCardsRepository(db *sql.DB) collections.Repository {
	return &collectionsRepo{db: db}
}

func (r *collectionsRepo) Create(ctx context.Context, newCollections *models.CollectionEntity) (*models.CollectionEntity, error) {
	collection := &models.CollectionEntity{}
	row := r.db.QueryRowContext(ctx, createCollection, &newCollections.CollectionName, &newCollections.UserID)
	err := row.Scan(collection)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.Create.QueryRowContext.Scan")
	}

	return collection, nil
}

func (r *collectionsRepo) GetByID(ctx context.Context, collectionID int) (*models.CollectionModel, error) {
	collection := &models.CollectionModel{}
	err := r.db.QueryRowContext(ctx, getCollectionById, collectionID).Scan(collection)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetByID.QueryRowContext.Scan")
	}

	return collection, nil
}

func (r *collectionsRepo) GetAllCollectionsByUserID(ctx context.Context, userID int, query *utilities.PaginationQuery) (*models.CollectionsList, error) {
	var totalRecords int
	err := r.db.QueryRowContext(ctx, getTotalCountByUserID, userID).Scan(&totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetAllCollectionsByUserID.QueryRowContext")
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
	rows, err := r.db.QueryContext(ctx, getCollectionsByUserID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetAllCollectionsByUserID.QueryContext")
	}
	for rows.Next() {
		var collection models.CollectionModel
		err := rows.Scan(&collection)
		if err != nil {
			return nil, errors.Wrap(err, "collectionsRepo.GetAllCollectionsByUserID.QueryContext.Scan")
		}
		collectionsList = append(collectionsList, &collection)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.GetAllCollectionsByUserID.rows.Err")
	}

	return &models.CollectionsList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         collectionsList,
	}, nil
}

func (r *collectionsRepo) Update(ctx context.Context, collectionToUpdate *models.CollectionEntity) (*models.CollectionEntity, error) {
	collection := &models.CollectionEntity{}
	err := r.db.QueryRowContext(ctx, updateCollection, &collectionToUpdate.CollectionName, &collectionToUpdate.CollectionID, &collectionToUpdate.UserID).Scan(collection)
	if err != nil {
		return nil, errors.Wrap(err, "collectionsRepo.Update.QueryRowContext.Scan")
	}

	return collection, nil
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

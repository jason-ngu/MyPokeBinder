package collections

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newCollections *models.CollectionEntity) (*models.CollectionEntity, error)
	GetByID(ctx context.Context, collectionID int) (*models.CollectionModel, error)
	GetAllCollectionsByUserID(ctx context.Context, userID int, query *utilities.PaginationQuery) (*models.CollectionsList, error)
	Update(ctx context.Context, collectionToUpdate *models.CollectionEntity) (*models.CollectionEntity, error)
	Delete(ctx context.Context, collectionId int) error
}

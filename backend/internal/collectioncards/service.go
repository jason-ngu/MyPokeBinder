package collectioncards

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	AddCardsToCollection(ctx context.Context, collectionID int, collectioncardsToAdd []*models.CollectionCardsModel) (*models.CollectionCardsList, error)
	GetByCollectionID(ctx context.Context, collectionID int, query *utilities.PaginationQuery) (*models.CollectionCardsList, error)
	RemoveCardsFromCollection(ctx context.Context, collectionID int, collectioncardsToRemove []*models.CollectionCardsModel) (*models.CollectionCardsList, error)
}

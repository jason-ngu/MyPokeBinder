package service

import (
	"backend/internal/collectioncards"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type collectioncardsService struct {
	collectioncardsRepo collectioncards.Repository
}

func NewCollectionCardsService(collectioncardsRepo collectioncards.Repository) collectioncards.Service {
	return &collectioncardsService{collectioncardsRepo: collectioncardsRepo}
}

func (s *collectioncardsService) AddCardsToCollection(ctx context.Context, collectionID int, collectioncardsToAdd []*models.CollectionCardsModel) (*models.CollectionCardsList, error) {
	return s.collectioncardsRepo.AddCardsToCollection(ctx, collectionID, collectioncardsToAdd)
}

func (s *collectioncardsService) GetByCollectionID(ctx context.Context, collectionID int, query *utilities.PaginationQuery) (*models.CollectionCardsList, error) {
	return s.collectioncardsRepo.GetByCollectionID(ctx, collectionID, query)
}

func (s *collectioncardsService) RemoveCardsFromCollection(ctx context.Context, collectionID int, collectioncardsToRemove []*models.CollectionCardsModel) (*models.CollectionCardsList, error) {
	return s.collectioncardsRepo.RemoveCardsFromCollection(ctx, collectionID, collectioncardsToRemove)
}

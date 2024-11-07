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

func (s *collectioncardsService) AddCardsToCollection(ctx context.Context, collectionId int, collectioncardsToAdd []*models.CollectionCardsModel) (*models.CollectionCardsList, error) {
	return s.collectioncardsRepo.AddCardsToCollection(ctx, collectionId, collectioncardsToAdd)
}

func (s *collectioncardsService) GetByCollectionID(ctx context.Context, collectionId int, query *utilities.PaginationQuery) (*models.CollectionCardsList, error) {
	return s.collectioncardsRepo.GetByCollectionID(ctx, collectionId, query)
}

func (s *collectioncardsService) RemoveCardsFromCollection(ctx context.Context, collectionId int, collectioncardsToRemove []*models.CollectionCardsModel) (*models.CollectionCardsList, error) {
	return s.collectioncardsRepo.RemoveCardsFromCollection(ctx, collectionId, collectioncardsToRemove)
}

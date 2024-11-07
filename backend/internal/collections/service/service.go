package service

import (
	"backend/internal/collections"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type collectionsService struct {
	collectionsRepo collections.Repository
}

func NewCollectionsService(collectionsRepo collections.Repository) collections.Service {
	return &collectionsService{collectionsRepo: collectionsRepo}
}

func (s *collectionsService) Create(ctx context.Context, newCollection *models.CollectionEntity) (*models.CollectionEntity, error) {
	return s.collectionsRepo.Create(ctx, newCollection)
}

func (s *collectionsService) GetByID(ctx context.Context, collectionId int) (*models.CollectionModel, error) {
	return s.collectionsRepo.GetByID(ctx, collectionId)
}

func (s *collectionsService) GetAllCollectionsByUserID(ctx context.Context, userID int, query *utilities.PaginationQuery) (*models.CollectionsList, error) {
	return s.collectionsRepo.GetAllCollectionsByUserID(ctx, userID, query)
}

func (s *collectionsService) Update(ctx context.Context, collectionToUpdate *models.CollectionEntity) (*models.CollectionEntity, error) {
	return s.collectionsRepo.Update(ctx, collectionToUpdate)
}

func (s *collectionsService) Delete(ctx context.Context, collectionId int) error {
	return s.collectionsRepo.Delete(ctx, collectionId)
}

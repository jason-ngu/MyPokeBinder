package cards

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newCard *models.CardEntity) (*models.CardEntity, error)
	GetByID(ctx context.Context, cardID int) (*models.CardModel, error)
	GetAllCards(ctx context.Context, searchParams models.CardSearchParams, query *utilities.PaginationQuery) (*models.CardsList, error)
	Update(ctx context.Context, cardToUpdate *models.CardEntity) (*models.CardEntity, error)
}

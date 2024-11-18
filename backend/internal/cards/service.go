package cards

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newCard *models.CardModel) (*models.CardModel, error)
	GetByID(ctx context.Context, cardID int) (*models.CardModel, error)
	Search(ctx context.Context, searchParams *models.CardSearchParams, query *utilities.PaginationQuery) (*models.CardsList, error)
	Update(ctx context.Context, cardID int, cardToUpdate *models.CardModel) (*models.CardModel, error)
}

package cardtypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newCardType *models.CardTypeEntity) (*models.CardTypeEntity, error)
	GetByCardID(ctx context.Context, cardID int, query *utilities.PaginationQuery) (*models.CardTypesList, error)
}

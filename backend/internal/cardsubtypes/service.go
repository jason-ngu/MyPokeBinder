package cardsubtypes

import (
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
)

type Service interface {
	Create(ctx context.Context, newCardSubtype *models.CardSubtypeEntity) (*models.CardSubtypeEntity, error)
	GetByCardID(ctx context.Context, cardID int, query *utilities.PaginationQuery) (*models.CardSubtypesList, error)
}

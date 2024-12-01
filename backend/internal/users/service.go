package users

import (
	"backend/internal/models"
	"context"
)

type Service interface {
	EnsureUser(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error)
	Create(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error)
	GetByID(ctx context.Context, userId int) (*models.UserModel, error)
	GetByProvider(ctx context.Context, providerType string, providerKey string) (*models.UserModel, error)
}

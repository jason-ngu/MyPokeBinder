package users

import (
	"backend/internal/models"
	"context"
)

type Repository interface {
	Create(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error)
	GetByID(ctx context.Context, userId int) (*models.UserModel, error)
}

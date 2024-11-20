package users

import (
	"backend/internal/models"
	"context"
)

type Service interface {
	CreateUser(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error)
	GetByID(ctx context.Context, userId int) (*models.UserModel, error)
}

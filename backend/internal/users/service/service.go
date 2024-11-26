package service

import (
	"backend/internal/models"
	"backend/internal/users"
	"context"
)

type usersService struct {
	usersRepo users.Repository
}

func NewUsersService(usersRepo users.Repository) users.Service {
	return &usersService{usersRepo: usersRepo}
}

func (s *usersService) Create(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error) {
	return s.usersRepo.Create(ctx, newUser)
}

func (s *usersService) GetByID(ctx context.Context, userId int) (*models.UserModel, error) {
	return s.usersRepo.GetByID(ctx, userId)
}

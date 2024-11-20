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

func (s *usersService) CreateUser(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error) {
	return s.usersRepo.CreateUser(ctx, newUser)
}

func (s *usersService) GetByID(ctx context.Context, userId int) (*models.UserModel, error) {
	return s.usersRepo.GetByID(ctx, userId)
}

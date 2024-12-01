package service

import (
	"backend/internal/models"
	"backend/internal/users"
	"context"
	"database/sql"
)

type usersService struct {
	usersRepo users.Repository
}

func NewUsersService(usersRepo users.Repository) users.Service {
	return &usersService{usersRepo: usersRepo}
}

func (s *usersService) EnsureUser(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error) {
	user, err := s.usersRepo.GetByProvider(ctx, newUser.ProviderType, newUser.ProviderKey)
	if err != nil {
		if err == sql.ErrNoRows {
			// Create new user if no results were returned
			return s.Create(ctx, newUser)
		}
		// If error was different return nil
		return nil, err
	}
	// If user was found return them
	return user, nil
}

func (s *usersService) Create(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error) {
	return s.usersRepo.Create(ctx, newUser)
}

func (s *usersService) GetByID(ctx context.Context, userId int) (*models.UserModel, error) {
	return s.usersRepo.GetByID(ctx, userId)
}

func (s *usersService) GetByProvider(ctx context.Context, providerType string, providerKey string) (*models.UserModel, error) {
	return s.usersRepo.GetByProvider(ctx, providerType, providerKey)
}

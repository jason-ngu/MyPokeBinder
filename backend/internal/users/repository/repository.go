package repository

import (
	"backend/internal/models"
	"backend/internal/users"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type usersRepo struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) users.Repository {
	return &usersRepo{db: db}
}

func (r *usersRepo) CreateUser(ctx context.Context, newUser *models.UserEntity) (*models.UserEntity, error) {
	user := &models.UserEntity{}
	row := r.db.QueryRowContext(ctx, createUser, &newUser.Name, &newUser.ProviderKey, &newUser.ProviderType)
	err := row.Scan(user)
	if err != nil {
		return nil, errors.Wrap(err, "usersRepo.Create.QueryRowContext.Scan")
	}

	return user, nil
}

func (r *usersRepo) GetByID(ctx context.Context, userId int) (*models.UserModel, error) {
	user := &models.UserModel{}
	err := r.db.QueryRowContext(ctx, getUserById, userId).Scan(user)
	if err != nil {
		return nil, errors.Wrap(err, "usersRepo.GetByID.QueryRowContext.Scan")
	}

	return user, nil
}

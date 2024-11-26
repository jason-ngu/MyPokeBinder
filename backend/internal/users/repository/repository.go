package repository

import (
	"backend/internal/models"
	"backend/internal/users"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type usersRepo struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) users.Repository {
	return &usersRepo{db: db}
}

func (r *usersRepo) Create(ctx context.Context, newUser *models.UserModel) (*models.UserModel, error) {
	user := &models.UserEntity{}
	row := r.db.QueryRowxContext(ctx, createUser, &newUser.Name, &newUser.ProviderKey, &newUser.ProviderType)
	err := row.StructScan(user)
	if err != nil {
		return nil, errors.Wrap(err, "usersRepo.Create.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, user.UserID)
}

func (r *usersRepo) GetByID(ctx context.Context, userId int) (*models.UserModel, error) {
	user := &models.UserModel{}
	row := r.db.QueryRowxContext(ctx, getUserById, userId)
	err := row.StructScan(user)
	if err != nil {
		return nil, errors.Wrap(err, "usersRepo.GetByID.QueryRowxContext.StructScan")
	}

	return user, nil
}

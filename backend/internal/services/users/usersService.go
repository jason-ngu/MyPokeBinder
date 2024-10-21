package usersService

import (
	internal "backend/internal"
	"backend/internal/models"
	usersRepository "backend/internal/repository/users"
)

func GetUser(env *internal.Env, user models.UserModel) (models.UserModel, error) {
	u, err := usersRepository.GetUser(env.DB, user)
	if err != nil {
		return models.UserModel{}, err
	}
	return u, nil
}

func CreateUser(env *internal.Env, newUser models.UserModel) (models.UserModel, error) {
	u, err := usersRepository.CreateUser(env.DB, newUser)
	if err != nil {
		return models.UserModel{}, err
	}
	return u, nil
}

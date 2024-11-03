package usersService

import (
	internal "backend/internal"
	"backend/internal/models"
	usersRepository "backend/internal/repository/users"
)

func GetUser(env *internal.Env, userId int) (models.UserModel, error) {
	u, err := usersRepository.GetUser(env.DB, userId)
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

func EnsureUser(env *internal.Env, user models.UserModel) (models.UserModel, error) {
	u, err := usersRepository.GetUserWithProviderInfo(env.DB, user.ProviderKey, user.ProviderType)
	if err != nil {
		return models.UserModel{}, err
	}
	// If there was no error, either user does not exist and needs to be created or can be returned
	if (u == models.UserModel{}) {
		newUser, err := usersRepository.CreateUser(env.DB, user)
		if err != nil {
			return models.UserModel{}, err
		}
		return newUser, nil
	} else {
		return u, nil
	}
}

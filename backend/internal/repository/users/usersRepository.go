package usersRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetUser(db *sql.DB, userId int) (models.UserModel, error) {
	row := db.QueryRow(`SELECT * FROM public.users
		WHERE user_id = $1`, userId)

	var u models.UserEntity
	err := row.Scan(&u.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserModel{}, nil
		}
		return models.UserModel{}, err
	}

	return models.UserModel{UserID: u.UserID}, nil
}

func GetUserWithProviderInfo(db *sql.DB, providerKey string, providerType string) (models.UserModel, error) {
	row := db.QueryRow(`SELECT * from public.users
		WHERE provider_key = $1 AND provider_type = $2`, providerKey, providerType)

	var u models.UserEntity
	err := row.Scan(&u.UserID, &u.Name, &u.ProviderKey, &u.ProviderType)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserModel{}, nil
		}
		return models.UserModel{}, err
	}

	return models.UserModel(u), nil
}

func CreateUser(db *sql.DB, newUser models.UserModel) (models.UserModel, error) {
	// existingUser, err := GetUser(db, newUser.UserID)
	// if (err == nil) && (existingUser != models.UserModel{}) {
	// 	return models.UserModel{}, errors.New("error creating user: user already exists")
	// } else if err != nil {
	// 	return models.UserModel{}, err
	// }

	_, err := db.Exec(`INSERT into public.users
		(name, provider_key, provider_type)
		VALUES ($1, $2, $3)`,
		newUser.Name, newUser.ProviderKey, newUser.ProviderType)
	if err != nil {
		return models.UserModel{}, err
	}

	user, err := GetUserWithProviderInfo(db, newUser.ProviderKey, newUser.ProviderType)
	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}

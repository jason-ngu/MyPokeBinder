package usersRepository

import (
	"backend/internal/models"
	"database/sql"
	"errors"
)

func GetUser(db *sql.DB, user models.UserModel) (models.UserModel, error) {
	row := db.QueryRow(`SELECT * FROM public.users 
		WHERE user_name = $1`, user.UserName)

	var u models.UserModel
	err := row.Scan(&u.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserModel{}, nil
		}
		return models.UserModel{}, err
	}

	return u, nil
}

func CreateUser(db *sql.DB, newUser models.UserModel) (models.UserModel, error) {
	existingUser, err := GetUser(db, newUser)
	if (err == nil) && (existingUser != models.UserModel{}) {
		return models.UserModel{}, errors.New("error creating user: user already exists")
	} else if err != nil {
		return models.UserModel{}, err
	}

	_, err = db.Exec(`INSERT into public.users
		(user_name)
		VALUES ($1)`,
		newUser.UserName)
	if err != nil {
		return models.UserModel{}, err
	}

	user, err := GetUser(db, newUser)
	if err != nil {
		return models.UserModel{}, err
	}

	return user, nil
}

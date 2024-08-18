package supertypesRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllSupertypes(db *sql.DB) ([]models.SupertypeModel, error) {
	rows, err := db.Query("SELECT * FROM public.supertypes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var supertypes []models.SupertypeModel

	for rows.Next() {
		var t models.SupertypeEntity

		err := rows.Scan(&t.SupertypeID, &t.SupertypeName)
		if err != nil {
			return nil, err
		}

		supertypes = append(supertypes, models.SupertypeModel(t))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return supertypes, nil
}

func GetSupertypeById(db *sql.DB, id int) (models.SupertypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.supertypes WHERE supertype_id = $1", id)

	var t models.SupertypeEntity
	err := row.Scan(&t.SupertypeID, &t.SupertypeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.SupertypeModel{}, nil
		}
		return models.SupertypeModel{}, err
	}

	return models.SupertypeModel(t), nil
}

func GetSupertypeByName(db *sql.DB, supertypeName string) (models.SupertypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.supertypes WHERE supertype_name = $1", supertypeName)

	var t models.SupertypeEntity
	err := row.Scan(&t.SupertypeID, &t.SupertypeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.SupertypeModel{}, nil
		}
		return models.SupertypeModel{}, err
	}

	return models.SupertypeModel(t), nil
}

func CreateSupertype(db *sql.DB, newSupertype models.SupertypeModel) (models.SupertypeModel, error) {
	supertype, err := GetSupertypeByName(db, newSupertype.SupertypeName)
	// If supertype already exists, return it
	if (err != nil) && (supertype != models.SupertypeModel{}) {
		return supertype, err
	}

	_, err = db.Exec("INSERT INTO public.supertypes (supertype_name) VALUES ($1)", newSupertype.SupertypeName)
	if err != nil {
		return models.SupertypeModel{}, err
	}

	createdSupertype, err := GetSupertypeByName(db, newSupertype.SupertypeName)
	if err != nil {
		return models.SupertypeModel{}, err
	}

	return createdSupertype, nil
}

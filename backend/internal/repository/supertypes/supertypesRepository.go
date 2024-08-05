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

		supertypes = append(supertypes, models.SupertypeModel{SupertypeName: t.SupertypeName})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return supertypes, nil
}

func GetSupertypeById(db *sql.DB, id int) (models.SupertypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.supertypes WHERE type_id = $1", id)
	if err := row.Err(); err != nil {
		return models.SupertypeModel{}, err
	}

	var t models.SupertypeEntity
	err := row.Scan(&t.SupertypeID, &t.SupertypeName)
	if err != nil {
		return models.SupertypeModel{}, err
	}

	return models.SupertypeModel{SupertypeName: t.SupertypeName}, nil
}

func GetSupertypeByName(db *sql.DB, supertypeName string) (models.SupertypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.supertypes WHERE supertype_name = $1", supertypeName)
	if err := row.Err(); err != nil {
		return models.SupertypeModel{}, err
	}

	var t models.SupertypeEntity
	err := row.Scan(&t.SupertypeID, &t.SupertypeName)
	if err != nil {
		return models.SupertypeModel{}, err
	}

	return models.SupertypeModel{SupertypeName: t.SupertypeName}, nil
}

func CreateSupertype(db *sql.DB, newSuperType models.SupertypeModel) (models.SupertypeModel, error) {
	supertype, err := GetSupertypeByName(db, newSuperType.SupertypeName)
	// If supertype already exists, return it
	if (err != nil) && (supertype != models.SupertypeModel{}) {
		return supertype, err
	}

	_, err = db.Exec("INSERT INTO public.supertypes (supertype_name) VALUES ($1)", newSuperType.SupertypeName)
	if err != nil {
		return models.SupertypeModel{}, err
	}

	newSupertype, err := GetSupertypeByName(db, newSuperType.SupertypeName)
	if err != nil {
		return models.SupertypeModel{}, err
	}

	return newSupertype, nil
}

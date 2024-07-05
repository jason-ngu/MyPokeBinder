package supertypesRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllSuperTypes(db *sql.DB) ([]models.SupertypeModel, error) {
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

func GetSupertypeByName(db *sql.DB, typeName string) (models.SupertypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.supertypes WHERE supertype_name = $1", typeName)
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

func CreateSupertype(db *sql.DB, newSuperType models.SupertypeModel) (int64, error) {
	_, err := GetSupertypeByName(db, newSuperType.SupertypeName)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO public.supertypes (supertype_name) VALUES ($1)", newSuperType.SupertypeName)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

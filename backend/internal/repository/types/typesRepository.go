package typesRepository

// https://gist.github.com/alexedwards/d42ae90aac9dfa75046ebf8a036b080b
// https://www.alexedwards.net/blog/organising-database-access

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllTypes(db *sql.DB) ([]models.TypeModel, error) {
	rows, err := db.Query("SELECT * FROM public.types")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var types []models.TypeModel

	for rows.Next() {
		var t models.TypeEntity

		err := rows.Scan(&t.TypeID, &t.TypeName)
		if err != nil {
			return nil, err
		}

		types = append(types, models.TypeModel{TypeName: t.TypeName})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return types, nil
}

func GetTypeById(db *sql.DB, id int) (models.TypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.types WHERE type_id = $1", id)
	if err := row.Err(); err != nil {
		return models.TypeModel{}, err
	}

	var t models.TypeEntity
	err := row.Scan(&t.TypeID, &t.TypeName)
	if err != nil {
		return models.TypeModel{}, err
	}

	return models.TypeModel{TypeName: t.TypeName}, nil
}

func GetTypeByName(db *sql.DB, typeName string) (models.TypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.types WHERE type_name = $1", typeName)
	if err := row.Err(); err != nil {
		return models.TypeModel{}, err
	}

	var t models.TypeEntity
	err := row.Scan(&t.TypeID, &t.TypeName)
	if err != nil {
		return models.TypeModel{}, err
	}

	return models.TypeModel{TypeName: t.TypeName}, nil
}

func CreateType(db *sql.DB, newType models.TypeModel) (models.TypeModel, error) {
	t, err := GetTypeByName(db, newType.TypeName)
	// If type already exists, return it
	if (err == nil) && (t != models.TypeModel{}) {
		return t, nil
	}

	// If it doesn't exist, try to create new type
	_, err = db.Exec("INSERT INTO public.types (type_name) VALUES ($1)", newType.TypeName)
	if err != nil {
		return models.TypeModel{}, err
	}

	newType, err = GetTypeByName(db, newType.TypeName)
	if err != nil {
		return models.TypeModel{}, err
	}

	return newType, nil
}

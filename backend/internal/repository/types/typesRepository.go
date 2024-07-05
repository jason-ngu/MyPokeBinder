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

// func createType(typeName string) {

// }

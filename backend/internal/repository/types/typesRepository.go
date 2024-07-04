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

// func getTypeById(id int) Types {

// }

// func createType(typeName string) {

// }

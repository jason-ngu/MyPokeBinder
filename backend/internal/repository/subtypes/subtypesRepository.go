package subtypesRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllSubTypes(db *sql.DB) ([]models.SubtypeModel, error) {
	rows, err := db.Query("SELECT * FROM public.subtypes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subtypes []models.SubtypeModel

	for rows.Next() {
		var t models.SubtypeEntity

		err := rows.Scan(&t.SubtypeID, &t.SubtypeName)
		if err != nil {
			return nil, err
		}

		subtypes = append(subtypes, models.SubtypeModel{SubtypeName: t.SubtypeName})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subtypes, nil
}

func GetSubtypeById(db *sql.DB, id int) (models.SubtypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.subtypes WHERE type_id = $1", id)
	if err := row.Err(); err != nil {
		return models.SubtypeModel{}, err
	}

	var t models.SubtypeEntity
	err := row.Scan(&t.SubtypeID, &t.SubtypeName)
	if err != nil {
		return models.SubtypeModel{}, err
	}

	return models.SubtypeModel{SubtypeName: t.SubtypeName}, nil
}

func GetSubtypeByName(db *sql.DB, subtypeName string) (models.SubtypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.subtypes WHERE subtype_name = $1", subtypeName)
	if err := row.Err(); err != nil {
		return models.SubtypeModel{}, err
	}

	var t models.SubtypeEntity
	err := row.Scan(&t.SubtypeID, &t.SubtypeName)
	if err != nil {
		return models.SubtypeModel{}, err
	}

	return models.SubtypeModel{SubtypeName: t.SubtypeName}, nil
}

func CreateSubtype(db *sql.DB, newSubType models.SubtypeModel) (int64, error) {
	_, err := GetSubtypeByName(db, newSubType.SubtypeName)
	if err != nil {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO public.subtypes (subtype_name) VALUES ($1)", newSubType.SubtypeName)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

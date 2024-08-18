package subtypesRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllSubtypes(db *sql.DB) ([]models.SubtypeModel, error) {
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

		subtypes = append(subtypes, models.SubtypeModel(t))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subtypes, nil
}

func GetSubtypeById(db *sql.DB, id int) (models.SubtypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.subtypes WHERE subtype_id = $1", id)

	var t models.SubtypeEntity
	err := row.Scan(&t.SubtypeID, &t.SubtypeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.SubtypeModel{}, nil
		}
		return models.SubtypeModel{}, err
	}

	return models.SubtypeModel(t), nil
}

func GetSubtypeByName(db *sql.DB, subtypeName string) (models.SubtypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.subtypes WHERE subtype_name = $1", subtypeName)

	var t models.SubtypeEntity
	err := row.Scan(&t.SubtypeID, &t.SubtypeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.SubtypeModel{}, nil
		}
		return models.SubtypeModel{}, err
	}

	return models.SubtypeModel(t), nil
}

func CreateSubtype(db *sql.DB, newSubtype models.SubtypeModel) (models.SubtypeModel, error) {
	subtype, err := GetSubtypeByName(db, newSubtype.SubtypeName)
	// If subtype already exists, return it
	if (err == nil) && (subtype != models.SubtypeModel{}) {
		return subtype, nil
	}

	_, err = db.Exec("INSERT INTO public.subtypes (subtype_name) VALUES ($1)", newSubtype.SubtypeName)
	if err != nil {
		return models.SubtypeModel{}, err
	}

	createdSubtype, err := GetSubtypeByName(db, newSubtype.SubtypeName)
	if err != nil {
		return models.SubtypeModel{}, err
	}

	return createdSubtype, nil
}

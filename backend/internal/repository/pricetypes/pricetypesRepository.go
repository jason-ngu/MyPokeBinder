package pricetypesRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllPricetypes(db *sql.DB) ([]models.PricetypeModel, error) {
	rows, err := db.Query("SELECT * FROM public.pricetypes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pricetypes []models.PricetypeModel

	for rows.Next() {
		var t models.PricetypeEntity

		err := rows.Scan(&t.PricetypeID, &t.PricetypeName)
		if err != nil {
			return nil, err
		}

		pricetypes = append(pricetypes, models.PricetypeModel(t))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pricetypes, nil
}

func GetPricetypeById(db *sql.DB, id int) (models.PricetypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.pricetypes WHERE pricetype_id = $1", id)

	var t models.PricetypeEntity
	err := row.Scan(&t.PricetypeID, &t.PricetypeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.PricetypeModel{}, nil
		}
		return models.PricetypeModel{}, err
	}

	return models.PricetypeModel(t), nil
}

func GetPricetypeByName(db *sql.DB, pricetypeName string) (models.PricetypeModel, error) {
	row := db.QueryRow("SELECT * FROM public.pricetypes WHERE pricetype_name = $1", pricetypeName)

	var t models.PricetypeEntity
	err := row.Scan(&t.PricetypeID, &t.PricetypeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.PricetypeModel{}, nil
		}
		return models.PricetypeModel{}, err
	}

	return models.PricetypeModel(t), nil
}

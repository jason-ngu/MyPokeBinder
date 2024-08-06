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

		pricetypes = append(pricetypes, models.PricetypeModel{PricetypeName: t.PricetypeName})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pricetypes, nil
}

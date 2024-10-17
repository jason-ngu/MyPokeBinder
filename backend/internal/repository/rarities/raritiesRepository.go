package raritiesRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllRarities(db *sql.DB) ([]models.RarityModel, error) {
	rows, err := db.Query("SELECT * FROM public.rarities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rarities []models.RarityModel

	for rows.Next() {
		var t models.RarityEntity

		err := rows.Scan(&t.RarityID, &t.RarityName)
		if err != nil {
			return nil, err
		}

		rarities = append(rarities, models.RarityModel(t))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rarities, nil
}

func GetRarityById(db *sql.DB, id int) (models.RarityModel, error) {
	row := db.QueryRow("SELECT * FROM public.rarities WHERE rarity_id = $1", id)

	var t models.RarityEntity
	err := row.Scan(&t.RarityID, &t.RarityName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.RarityModel{}, nil
		}
		return models.RarityModel{}, err
	}

	return models.RarityModel(t), nil
}

func GetRarityByName(db *sql.DB, raritiesName string) (models.RarityModel, error) {
	row := db.QueryRow("SELECT * FROM public.rarities WHERE rarity_name = $1", raritiesName)

	var t models.RarityEntity
	err := row.Scan(&t.RarityID, &t.RarityName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.RarityModel{}, nil
		}
		return models.RarityModel{}, err
	}

	return models.RarityModel(t), nil
}

func CreateRarity(db *sql.DB, newRarity models.RarityModel) (models.RarityModel, error) {
	rarity, err := GetRarityByName(db, newRarity.RarityName)
	// If rarity already exists, return it
	if (err == nil) && (rarity != models.RarityModel{}) {
		return rarity, err
	}

	_, err = db.Exec("INSERT INTO public.rarities (rarity_name) VALUES ($1)", newRarity.RarityName)
	if err != nil {
		return models.RarityModel{}, err
	}

	createdRarity, err := GetRarityByName(db, newRarity.RarityName)
	if err != nil {
		return models.RarityModel{}, err
	}

	return createdRarity, nil
}

package raritiesService

import (
	internal "backend/internal"
	"backend/internal/models"
	raritiesRepository "backend/internal/repository/rarities"
)

func GetAllRarities(env *internal.Env) ([]models.RarityModel, error) {
	allRarities, err := raritiesRepository.GetAllRarities(env.DB)
	if err != nil {
		return []models.RarityModel{}, err
	}
	return allRarities, nil
}

func GetRarityById(env *internal.Env, id int) (models.RarityModel, error) {
	rarity, err := raritiesRepository.GetRarityById(env.DB, id)
	if err != nil {
		return models.RarityModel{}, err
	}
	return rarity, nil
}

func CreateRarity(env *internal.Env, newRarity models.RarityModel) (models.RarityModel, error) {
	rarity, err := raritiesRepository.CreateRarity(env.DB, newRarity)
	if err != nil {
		return models.RarityModel{}, err
	}
	return rarity, nil
}

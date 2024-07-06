package raritiesService

import (
	internal "backend/internal"
	"backend/internal/models"
	raritiesRepository "backend/internal/repository/rarities"
)

func GetAllRarities(env *internal.Env) ([]models.RarityModel, error) {
	allTypes, err := raritiesRepository.GetAllRarities(env.DB)
	if err != nil {
		return []models.RarityModel{}, err
	}
	return allTypes, nil
}

func GetRarityById(env *internal.Env, id int) (models.RarityModel, error) {
	t, err := raritiesRepository.GetRarityById(env.DB, id)
	if err != nil {
		return models.RarityModel{}, err
	}
	return t, nil
}

func CreateRarity(env *internal.Env, newRarity models.RarityModel) (int64, error) {
	id, err := raritiesRepository.CreateRarity(env.DB, newRarity)
	if err != nil {
		return 0, err
	}
	return id, nil
}

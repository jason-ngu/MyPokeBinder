package setsService

import (
	internal "backend/internal"
	"backend/internal/models"
	setsRepository "backend/internal/repository/sets"
)

func GetAllSets(env *internal.Env) ([]models.SetModel, error) {
	allSets, err := setsRepository.GetAllSets(env.DB)
	if err != nil {
		return []models.SetModel{}, err
	}
	return allSets, nil
}

func GetSetById(env *internal.Env, id int) (models.SetModel, error) {
	set, err := setsRepository.GetSetById(env.DB, id)
	if err != nil {
		return models.SetModel{}, err
	}
	return set, nil
}

func GetSetByName(env *internal.Env, setName string) (models.SetModel, error) {
	set, err := setsRepository.GetSetByName(env.DB, setName)
	if err != nil {
		return models.SetModel{}, err
	}
	return set, nil
}

func CreateSet(env *internal.Env, newSet models.SetModel) (models.SetModel, error) {
	set, err := setsRepository.CreateSet(env.DB, newSet)
	if err != nil {
		return models.SetModel{}, err
	}
	return set, nil
}

func UpdateSet(env *internal.Env, setToUpdate models.SetModel) (models.SetModel, error) {
	set, err := setsRepository.UpdateSet(env.DB, setToUpdate)
	if err != nil {
		return models.SetModel{}, err
	}
	return set, nil
}

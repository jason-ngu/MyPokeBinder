package subtypesService

import (
	internal "backend/internal"
	"backend/internal/models"
	subtypesRepository "backend/internal/repository/subtypes"
)

func GetAllSubtypes(env *internal.Env) ([]models.SubtypeModel, error) {
	allSubtypes, err := subtypesRepository.GetAllSubtypes(env.DB)
	if err != nil {
		return []models.SubtypeModel{}, err
	}
	return allSubtypes, nil
}

func GetSubtypeById(env *internal.Env, id int) (models.SubtypeModel, error) {
	subtype, err := subtypesRepository.GetSubtypeById(env.DB, id)
	if err != nil {
		return models.SubtypeModel{}, err
	}
	return subtype, nil
}

func CreateSubtype(env *internal.Env, newSubtype models.SubtypeModel) (models.SubtypeModel, error) {
	subtype, err := subtypesRepository.CreateSubtype(env.DB, newSubtype)
	if err != nil {
		return models.SubtypeModel{}, err
	}
	return subtype, nil
}

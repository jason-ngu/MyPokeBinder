package subtypesService

import (
	internal "backend/internal"
	"backend/internal/models"
	subtypesRepository "backend/internal/repository/subtypes"
)

func GetAllSubTypes(env *internal.Env) ([]models.SubtypeModel, error) {
	allTypes, err := subtypesRepository.GetAllSubTypes(env.DB)
	if err != nil {
		return []models.SubtypeModel{}, err
	}
	return allTypes, nil
}

func GetSubtypeById(env *internal.Env, id int) (models.SubtypeModel, error) {
	t, err := subtypesRepository.GetSubtypeById(env.DB, id)
	if err != nil {
		return models.SubtypeModel{}, err
	}
	return t, nil
}

func CreateSubtype(env *internal.Env, newSubtype models.SubtypeModel) (int64, error) {
	id, err := subtypesRepository.CreateSubtype(env.DB, newSubtype)
	if err != nil {
		return 0, err
	}
	return id, nil
}

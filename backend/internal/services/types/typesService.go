package typesService

import (
	internal "backend/internal"
	"backend/internal/models"
	typesRepository "backend/internal/repository/types"
)

func GetAllTypes(env *internal.Env) ([]models.TypeModel, error) {
	allTypes, err := typesRepository.GetAllTypes(env.DB)
	if err != nil {
		return nil, err
	}
	return allTypes, nil
}

func GetTypeById(env *internal.Env, id int) (models.TypeModel, error) {
	t, err := typesRepository.GetTypeById(env.DB, id)
	if err != nil {
		return models.TypeModel{}, err
	}
	return t, nil
}

func CreateType(env *internal.Env, newType models.TypeModel) (int64, error) {
	id, err := typesRepository.CreateType(env.DB, newType)
	if err != nil {
		return 0, err
	}
	return id, nil
}

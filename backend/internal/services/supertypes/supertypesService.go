package supertypesService

import (
	internal "backend/internal"
	"backend/internal/models"
	supertypesRepository "backend/internal/repository/supertypes"
)

func GetAllSupertypes(env *internal.Env) ([]models.SupertypeModel, error) {
	allTypes, err := supertypesRepository.GetAllSupertypes(env.DB)
	if err != nil {
		return []models.SupertypeModel{}, err
	}
	return allTypes, nil
}

func GetSupertypeById(env *internal.Env, id int) (models.SupertypeModel, error) {
	t, err := supertypesRepository.GetSupertypeById(env.DB, id)
	if err != nil {
		return models.SupertypeModel{}, err
	}
	return t, nil
}

func CreateSupertype(env *internal.Env, newSuperType models.SupertypeModel) (models.SupertypeModel, error) {
	supertype, err := supertypesRepository.CreateSupertype(env.DB, newSuperType)
	if err != nil {
		return models.SupertypeModel{}, err
	}
	return supertype, nil
}

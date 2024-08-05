package supertypesService

import (
	internal "backend/internal"
	"backend/internal/models"
	supertypesRepository "backend/internal/repository/supertypes"
)

func GetAllSupertypes(env *internal.Env) ([]models.SupertypeModel, error) {
	allSupertypes, err := supertypesRepository.GetAllSupertypes(env.DB)
	if err != nil {
		return []models.SupertypeModel{}, err
	}
	return allSupertypes, nil
}

func GetSupertypeById(env *internal.Env, id int) (models.SupertypeModel, error) {
	supertype, err := supertypesRepository.GetSupertypeById(env.DB, id)
	if err != nil {
		return models.SupertypeModel{}, err
	}
	return supertype, nil
}

func CreateSupertype(env *internal.Env, newSuperType models.SupertypeModel) (models.SupertypeModel, error) {
	supertype, err := supertypesRepository.CreateSupertype(env.DB, newSuperType)
	if err != nil {
		return models.SupertypeModel{}, err
	}
	return supertype, nil
}

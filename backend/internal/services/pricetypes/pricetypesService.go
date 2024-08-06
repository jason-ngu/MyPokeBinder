package pricetypesService

import (
	internal "backend/internal"
	"backend/internal/models"
	pricetypesRepository "backend/internal/repository/pricetypes"
)

func GetAllPricetypes(env *internal.Env) ([]models.PricetypeModel, error) {
	allPricetypes, err := pricetypesRepository.GetAllPricetypes(env.DB)
	if err != nil {
		return []models.PricetypeModel{}, err
	}
	return allPricetypes, nil
}

package typesService

import (
	"backend/internal/models"
	typesRepository "backend/internal/repository/types"
	"backend/internal/services"
)

type Types struct {
	TypeName string
}

func GetAllTypes(env *services.Env) []models.TypesModel {
	allTypes := typesRepository.GetAllTypes(env.DB)
	return allTypes
}

// func getTypeById(id int) Types {

// }

// func createType(typeName string) {

// }

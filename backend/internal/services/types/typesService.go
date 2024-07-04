package typesService

import (
	internal "backend/internal"
	"backend/internal/models"
	typesRepository "backend/internal/repository/types"
)

type Types struct {
	TypeName string
}

func GetAllTypes(env *internal.Env) []models.TypesModel {
	allTypes := typesRepository.GetAllTypes(env.DB)
	return allTypes
}

// func getTypeById(id int) Types {

// }

// func createType(typeName string) {

// }

package typesService

import (
	internal "backend/internal"
	"backend/internal/models"
	typesRepository "backend/internal/repository/types"
)

type Types struct {
	TypeName string
}

func GetAllTypes(env *internal.Env) ([]models.TypeModel, error) {
	allTypes, err := typesRepository.GetAllTypes(env.DB)
	if err != nil {
		return nil, err
	}
	return allTypes, nil
}

// func getTypeById(id int) Types {

// }

// func createType(typeName string) {

// }

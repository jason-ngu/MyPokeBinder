package collectionsService

import (
	internal "backend/internal"
	"backend/internal/models"
	collectionsRepository "backend/internal/repository/collections"
)

func GetUserCollections(env *internal.Env, userId int) ([]models.CollectionModel, error) {
	collections, err := collectionsRepository.GetUserCollections(env.DB, userId)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func CreateCollection(env *internal.Env, collectionName string, userId int) (models.CollectionModel, error) {
	newCollection, err := collectionsRepository.CreateCollection(env.DB, collectionName, userId)
	if err != nil {
		return models.CollectionModel{}, err
	}
	return newCollection, nil
}

func UpdateCollection(env *internal.Env, collectionToUpdate models.CollectionModel) (models.CollectionModel, error) {
	updatedCollection, err := collectionsRepository.UpdateCollection(env.DB, collectionToUpdate)
	if err != nil {
		return models.CollectionModel{}, err
	}
	return updatedCollection, nil
}

func DeleteCollection(env *internal.Env, collectionToDelete models.CollectionModel) (bool, error) {
	_, err := collectionsRepository.DeleteCollection(env.DB, collectionToDelete)
	if err != nil {
		return false, err
	}
	return true, nil
}

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

func GetCollection(env *internal.Env, collectionId int) (models.CollectionModel, error) {
	collection, err := collectionsRepository.GetCollection(env.DB, collectionId)
	if err != nil {
		return models.CollectionModel{}, err
	}
	return collection, nil
}

func CreateCollection(env *internal.Env, collectionToCreate models.CollectionModel) (models.CollectionModel, error) {
	newCollection, err := collectionsRepository.CreateCollection(env.DB, collectionToCreate)
	if err != nil {
		return models.CollectionModel{}, err
	}
	return newCollection, nil
}

func UpdateCollection(env *internal.Env, collectionId int, collectionToUpdate models.CollectionModel) (models.CollectionModel, error) {
	updatedCollection, err := collectionsRepository.UpdateCollection(env.DB, collectionId, collectionToUpdate)
	if err != nil {
		return models.CollectionModel{}, err
	}
	return updatedCollection, nil
}

func DeleteCollection(env *internal.Env, collectionId int) (bool, error) {
	_, err := collectionsRepository.DeleteCollection(env.DB, collectionId)
	if err != nil {
		return false, err
	}
	return true, nil
}

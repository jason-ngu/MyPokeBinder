package collectioncards

import (
	internal "backend/internal"
	"backend/internal/models"
	collectioncardsRepository "backend/internal/repository/collectioncards"
)

// func GetUserCollectionCards(env *internal.Env, userId int) ([]models.CollectionCardsModel, error) {
// 	userCollections, err := collectioncardsRepository.GetUserCollectionCards(env.DB, userId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return userCollections, nil
// }

func GetCollectionCards(env *internal.Env, collectionId int) ([]models.CollectionCardsModel, error) {
	collectionCards, err := collectioncardsRepository.GetCollectionCards(env.DB, collectionId)
	if err != nil {
		return nil, err
	}
	return collectionCards, nil
}

func AddCardsToCollection(env *internal.Env, collectionId int, collectioncardsToAdd []models.CollectionCardsModel) ([]models.CollectionCardsModel, error) {
	collectionCards, err := collectioncardsRepository.AddCardsToCollection(env.DB, collectionId, collectioncardsToAdd)
	if err != nil {
		return nil, err
	}
	return collectionCards, nil
}

func RemoveCardsFromCollection(env *internal.Env, collectionId int, collectioncardsToRemove []models.CollectionCardsModel) ([]models.CollectionCardsModel, error) {
	collectionCards, err := collectioncardsRepository.RemoveCardsFromCollection(env.DB, collectionId, collectioncardsToRemove)
	if err != nil {
		return nil, err
	}
	return collectionCards, nil
}

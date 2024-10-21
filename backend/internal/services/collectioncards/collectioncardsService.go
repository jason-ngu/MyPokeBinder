package collectioncards

import (
	internal "backend/internal"
	"backend/internal/models"
	collectioncardsRepository "backend/internal/repository/collectioncards"
)

func GetUserCollectionCards(env *internal.Env, userId int) ([]models.CollectionCardsModel, error) {
	userCollections, err := collectioncardsRepository.GetUserCollectionCards(env.DB, userId)
	if err != nil {
		return nil, err
	}
	return userCollections, nil
}

func GetCollectionCards(env *internal.Env, collectionId int) (models.CollectionCardsModel, error) {
	collectionCards, err := collectioncardsRepository.GetCollectionCards(env.DB, collectionId)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}
	return collectionCards, nil
}

func AddCardsToCollection(env *internal.Env, collectionId int, cardIds []int) (models.CollectionCardsModel, error) {
	collectionCards, err := collectioncardsRepository.AddCardsToCollection(env.DB, collectionId, cardIds)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}
	return collectionCards, nil
}

func RemoveCardsFromCollection(env *internal.Env, collectionId int, cardIds []int) (models.CollectionCardsModel, error) {
	collectionCards, err := collectioncardsRepository.RemoveCardsFromCollection(env.DB, collectionId, cardIds)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}
	return collectionCards, nil
}

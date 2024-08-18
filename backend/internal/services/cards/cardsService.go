package cardsService

import (
	internal "backend/internal"
	"backend/internal/models"
	cardsRepository "backend/internal/repository/cards"
)

func GetAllCards(env *internal.Env) ([]models.CardModel, error) {
	allCards, err := cardsRepository.GetAllCards(env.DB)
	if err != nil {
		return []models.CardModel{}, err
	}
	return allCards, nil
}

func SearchCards(env *internal.Env, searchParams models.CardSearchParams) ([]models.CardModel, error) {
	allCards, err := cardsRepository.SearchCards(env.DB, searchParams)
	if err != nil {
		return []models.CardModel{}, err
	}
	return allCards, nil
}

func GetCard(env *internal.Env, cardCode string, pricetype string) (models.CardModel, error) {
	card, err := cardsRepository.GetCard(env.DB, cardCode, pricetype)
	if err != nil {
		return models.CardModel{}, err
	}
	return card, nil
}

func CreateCard(env *internal.Env, newCard models.CardModel) (models.CardModel, error) {
	card, err := cardsRepository.CreateCard(env.DB, newCard)
	if err != nil {
		return models.CardModel{}, err
	}
	return card, nil
}

func UpdateCard(env *internal.Env, cardToUpdate models.CardModel) (models.CardModel, error) {
	card, err := cardsRepository.UpdateCard(env.DB, cardToUpdate)
	if err != nil {
		return models.CardModel{}, err
	}
	return card, nil
}

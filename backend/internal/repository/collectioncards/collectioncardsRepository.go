package collectioncardsRepository

import (
	"backend/internal/models"
	cardsRepository "backend/internal/repository/cards"
	collectionsRepository "backend/internal/repository/collections"
	"database/sql"
	"fmt"
	"strings"
)

func GetUserCollectionCards(db *sql.DB, userId int) ([]models.CollectionCardsModel, error) {
	userCollections, err := collectionsRepository.GetUserCollections(db, userId)
	if err != nil {
		return nil, err
	}

	var userCollectionCards []models.CollectionCardsModel

	for _, collection := range userCollections {
		collectionCard, err := GetCollectionCards(db, collection.CollectionID)
		if err != nil {
			return nil, err
		}
		userCollectionCards = append(userCollectionCards, collectionCard)
	}

	return userCollectionCards, nil
}

func GetCollectionCards(db *sql.DB, collectionId int) (models.CollectionCardsModel, error) {
	rows, err := db.Query("SELECT * FROM public.collectioncards WHERE collection_id = $1", collectionId)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}
	defer rows.Close()

	var cardIds []int

	for rows.Next() {
		var t models.CollectionCardsEntity

		err := rows.Scan(&t.CollectionID, &t.CardID)
		if err != nil {
			return models.CollectionCardsModel{}, err
		}

		cardIds = append(cardIds, t.CardID)
	}
	if err = rows.Err(); err != nil {
		return models.CollectionCardsModel{}, err
	}

	cardModels, err := cardsRepository.GetCardsById(db, cardIds)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}

	collectionName, err := collectionsRepository.GetCollectionNameById(db, collectionId)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}

	var collectionCards = models.CollectionCardsModel{
		CollectionID:    collectionId,
		CollectionName:  collectionName,
		CollectionCards: cardModels,
	}

	return collectionCards, nil
}

func AddCardsToCollection(db *sql.DB, collectionId int, cardIds []int) (models.CollectionCardsModel, error) {
	// Check if collection exists first
	_, err := collectionsRepository.GetCollectionNameById(db, collectionId)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}

	var strValues []string
	for _, id := range cardIds {
		strValues = append(strValues, fmt.Sprintf("(%d, %d)", collectionId, id))
	}

	// Join the string slice with commas
	cardIdsStr := strings.Join(strValues, ", ")

	// Can allow multiple of the same card in a collection
	_, err = db.Exec(`INSERT INTO public.collectioncards
		(collection_id, card_id)
		VALUES $1`, cardIdsStr)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}

	collectionCards, err := GetCollectionCards(db, collectionId)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}

	return collectionCards, nil
}

func RemoveCardsFromCollection(db *sql.DB, collectionId int, cardIds []int) (models.CollectionCardsModel, error) {
	// Check if collection exists first
	_, err := collectionsRepository.GetCollectionNameById(db, collectionId)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}

	var strNumbers []string
	for _, id := range cardIds {
		strNumbers = append(strNumbers, fmt.Sprintf("%d", id))
	}

	// Join the string slice with commas
	cardIdsStr := strings.Join(strNumbers, ", ")

	// Can allow multiple of the same card in a collection
	_, err = db.Exec(`DELETE FROM public.collectioncards
		WHERE collection_id = $1 AND card_id IN $2`, collectionId, cardIdsStr)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}

	collectionCards, err := GetCollectionCards(db, collectionId)
	if err != nil {
		return models.CollectionCardsModel{}, err
	}

	return collectionCards, nil
}

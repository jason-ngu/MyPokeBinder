package collectioncardsRepository

import (
	"backend/internal/models"
	cardsRepository "backend/internal/repository/cards"
	collectionsRepository "backend/internal/repository/collections"
	"database/sql"
	"fmt"
)

// func GetUserCollectionCards(db *sql.DB, userId int) ([]models.CollectionCardsModel, error) {
// 	userCollections, err := collectionsRepository.GetUserCollections(db, userId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var userCollectionCards []models.CollectionCardsModel

// 	for _, collection := range userCollections {
// 		collectionCard, err := GetCollectionCards(db, collection.CollectionID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		userCollectionCards = append(userCollectionCards, collectionCard)
// 	}

// 	return userCollectionCards, nil
// }

func GetCollectionCards(db *sql.DB, collectionId int) ([]models.CollectionCardsModel, error) {
	rows, err := db.Query("SELECT * FROM public.collectioncards WHERE collection_id = $1", collectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collectionCardsSlice []models.CollectionCardsModel

	for rows.Next() {
		var t models.CollectionCardsEntity

		err := rows.Scan(&t.CollectionID, &t.CardID, &t.Quantity, &t.Grade, &t.GradingCompany)
		if err != nil {
			return nil, err
		}

		card, err := cardsRepository.GetCardsById(db, []int{t.CardID})
		if err != nil {
			return nil, err
		}

		collectionCardsSlice = append(collectionCardsSlice, models.CollectionCardsModel{
			CollectionID:   t.CollectionID,
			Card:           card[0],
			Quantity:       t.Quantity,
			Grade:          t.Grade,
			GradingCompany: t.GradingCompany,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return collectionCardsSlice, nil
}

func AddCardsToCollection(db *sql.DB, collectionId int, collectioncardsToAdd []models.CollectionCardsModel) ([]models.CollectionCardsModel, error) {
	// Check if collection exists first
	_, err := collectionsRepository.GetCollection(db, collectionId)
	if err != nil {
		return nil, err
	}

	var strValues []string
	for _, collectioncard := range collectioncardsToAdd {
		strValues = append(strValues, fmt.Sprintf("(%d, %d, %d, %d, %s)",
			collectionId, collectioncard.Card.CardID, collectioncard.Quantity, collectioncard.Grade, collectioncard.GradingCompany))

		_, err = db.Exec(`INSERT INTO collectioncards (collection_id, card_id, quantity, grade, grading_company)
			VALUES $1
			ON CONFLICT (collection_id, card_id, grade, grading_company)
			DO UPDATE SET quantity = collectioncards.quantity + $2;`, strValues, collectioncard.Quantity)
		if err != nil {
			return nil, err
		}
	}

	// Join the string slice with commas
	// cardIdsStr := strings.Join(strValues, ", ")

	// Can allow multiple of the same card in a collection
	// _, err = db.Exec(`INSERT INTO public.collectioncards
	// 	(collection_id, card_id, quantity, grade, grading_company)
	// 	VALUES $1`, cardIdsStr)
	// if err != nil {
	// 	return nil, err
	// }

	collectionCards, err := GetCollectionCards(db, collectionId)
	if err != nil {
		return nil, err
	}

	return collectionCards, nil
}

func RemoveCardsFromCollection(db *sql.DB, collectionId int, collectioncardsToRemove []models.CollectionCardsModel) ([]models.CollectionCardsModel, error) {
	// Check if collection exists first
	_, err := collectionsRepository.GetCollection(db, collectionId)
	if err != nil {
		return nil, err
	}

	// var strNumbers []string
	for _, collectioncard := range collectioncardsToRemove {
		_, err = db.Exec(`WITH updated AS (
				UPDATE collectioncards
				SET quantity = GREATEST(quantity - $1, 0)
				WHERE collection_id = $2 AND card_id = $3 AND grade = $4 AND grading_company = $5
				RETURNING *
			)
			DELETE FROM collectioncards
			WHERE collection_id = $2 AND card_id = $3 AND quantity = 0;`,
			collectioncard.Quantity, collectioncard.Card.CardID, collectioncard.Grade, collectioncard.GradingCompany)
		if err != nil {
			return nil, err
		}
	}

	// // Join the string slice with commas
	// cardIdsStr := strings.Join(strNumbers, ", ")

	// // Can allow multiple of the same card in a collection
	// _, err = db.Exec(`DELETE FROM public.collectioncards
	// 	WHERE collection_id = $1 AND card_id IN $2`, collectionId, cardIdsStr)
	// if err != nil {
	// 	return nil, err
	// }

	collectionCards, err := GetCollectionCards(db, collectionId)
	if err != nil {
		return nil, err
	}

	return collectionCards, nil
}

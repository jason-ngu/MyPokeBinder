package repository

const (
	addCardsToCollection = `INSERT INTO collectioncards (collection_id, card_id, quantity, grade, grading_company)
			VALUES $1
			ON CONFLICT (collection_id, card_id, grade, grading_company)
			DO UPDATE SET quantity = collectioncards.quantity + $2;`
	getByCollectionID           = `SELECT * FROM public.collectioncards WHERE collection_id = $1`
	getTotalCountByCollectionID = `SELECT COUNT(*) FROM public.collectioncards WHERE collection_id = $1`
	removeCardsFromCollection   = `WITH updated AS (
				UPDATE collectioncards
				SET quantity = GREATEST(quantity - $1, 0)
				WHERE collection_id = $2 AND card_id = $3 AND grade = $4 AND grading_company = $5
				RETURNING *
			)
			DELETE FROM collectioncards
			WHERE collection_id = $2 AND card_id = $3 AND quantity = 0;`
)

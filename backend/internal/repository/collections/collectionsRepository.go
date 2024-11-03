package collectionsRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetUserCollections(db *sql.DB, userId int) ([]models.CollectionModel, error) {
	rows, err := db.Query("SELECT * FROM public.collections WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collections []models.CollectionModel

	for rows.Next() {
		var t models.CollectionEntity

		err := rows.Scan(&t.CollectionID, &t.CollectionName, &t.UserID)
		if err != nil {
			return nil, err
		}

		collections = append(collections, models.CollectionModel{
			CollectionID:   t.CollectionID,
			CollectionName: t.CollectionName,
			UserID:         userId,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return collections, nil
}

func GetCollection(db *sql.DB, collectionId int) (models.CollectionModel, error) {
	collection := db.QueryRow("SELECT * FROM public.collections WHERE collection_id = $1", collectionId)

	var collectionEntity models.CollectionEntity
	err := collection.Scan(&collectionEntity.CollectionID, &collectionEntity.CollectionName, &collectionEntity.UserID)
	if err != nil {
		return models.CollectionModel{}, err
	}

	return models.CollectionModel(collectionEntity), nil
}

func CreateCollection(db *sql.DB, collectionToCreate models.CollectionModel) (models.CollectionModel, error) {
	result, err := db.Exec(`INSERT INTO public.collections
		(collection_name, user_id)
		VALUES ($1, $2)`, collectionToCreate.CollectionName, collectionToCreate.UserID)
	if err != nil {
		return models.CollectionModel{}, err
	}

	newCollectionId, err := result.LastInsertId()
	if err != nil {
		return models.CollectionModel{}, err
	}

	return models.CollectionModel{
		CollectionID:   int(newCollectionId),
		CollectionName: collectionToCreate.CollectionName,
		UserID:         collectionToCreate.UserID,
	}, nil
}

func UpdateCollection(db *sql.DB, collectionId int, collectionToUpdate models.CollectionModel) (models.CollectionModel, error) {
	_, err := db.Exec(`UPDATE public.collections
		SET collection_name = $1
		WHERE collection_id = $2 AND user_id = $3`,
		collectionToUpdate.CollectionName, collectionId, collectionToUpdate.UserID)
	if err != nil {
		return models.CollectionModel{}, err
	}

	return collectionToUpdate, nil
}

func DeleteCollection(db *sql.DB, collectionId int) (bool, error) {
	_, err := db.Exec(`DELETE FROM public.collections
		WHERE collection_id = $1`, collectionId)
	if err != nil {
		return false, err
	}

	return true, nil
}

package models

type CollectionEntity struct {
	CollectionID   int    `json:"collection_id"`
	CollectionName string `json:"collection_name"`
	UserID         int    `json:"user_id"`
}

type CollectionModel struct {
	CollectionID   int
	CollectionName string
	UserID         int
}

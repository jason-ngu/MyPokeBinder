package models

type CollectionCardsEntity struct {
	CollectionID int `json:"collection_id"`
	CardID       int `json:"card_id"`
}

type CollectionCardsModel struct {
	CollectionID    int
	CollectionName  string
	CollectionCards []CardModel
}

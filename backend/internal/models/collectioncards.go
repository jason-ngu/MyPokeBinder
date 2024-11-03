package models

type CollectionCardsEntity struct {
	CollectionID   int    `json:"collection_id"`
	CardID         int    `json:"card_id"`
	Quantity       int    `json:"quantity"`
	Grade          int    `json:"grade"`
	GradingCompany string `json:"grading_company"`
}

type CollectionCardsModel struct {
	CollectionID   int
	Card           CardModel
	Quantity       int
	Grade          int
	GradingCompany string
	// CollectionCards []CardModel
}

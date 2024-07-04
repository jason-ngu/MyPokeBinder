package models

type RarityEntity struct {
	RarityID   int    `json:"rarity_id"`
	RarityName string `json:"rarity_name"`
}

type RarityModel struct {
	RarityName string
}

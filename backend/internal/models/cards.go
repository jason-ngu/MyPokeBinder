package models

type CardEntity struct {
	CardId          int     `json:"card_id"`
	ApiId           string  `json:"api_id"`
	CardName        string  `json:"card_name"`
	SetId           int     `json:"set_id"`
	SupertypeId     int     `json:"supertype_id"`
	RarityId        int     `json:"rarity_id"`
	MarketPrice     float32 `json:"market_price"`
	PricetypeId     int     `json:"pricetype_id"`
	Image           string  `json:"image"`
	SyncDateCreated string  `json:"sync_date_created"`
	SyncDateUpdated string  `json:"sync_date_updated"`
}

type CardModel struct {
	CardName      string
	SetName       string
	SupertypeName string
	RarityName    string
	MarketPrice   float32
	PricetypeName string
	Image         string
}

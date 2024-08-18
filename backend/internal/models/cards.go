package models

type CardEntity struct {
	CardID          int     `json:"card_id"`
	CardCode        string  `json:"card_code"`
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
	CardID        int
	CardName      string
	CardCode      string
	SetName       string
	SupertypeName string
	RarityName    string
	MarketPrice   float32
	PricetypeName string
	Image         string
}

type CardSearchParams struct {
	CardName      string
	CardCode      string
	SetName       string
	SupertypeName string
	RarityName    string
	PricetypeName string
}

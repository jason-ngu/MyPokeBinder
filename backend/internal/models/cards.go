package models

type CardEntity struct {
	CardID          int     `json:"card_id" db:"card_id"`
	CardCode        string  `json:"card_code" db:"card_code"`
	CardName        string  `json:"card_name" db:"card_name"`
	SetID           int     `json:"set_id" db:"set_id"`
	SupertypeID     int     `json:"supertype_id" db:"supertype_id"`
	RarityID        int     `json:"rarity_id" db:"rarity_id"`
	MarketPrice     float32 `json:"market_price" db:"market_price"`
	PricetypeID     int     `json:"pricetype_id" db:"pricetype_id"`
	Image           string  `json:"image" db:"image"`
	SyncDateCreated string  `json:"sync_date_created" db:"sync_date_created"`
	SyncDateUpdated string  `json:"sync_date_updated" db:"sync_date_updated"`
}

type CardModel struct {
	CardID      int            `json:"card_id" db:"card_id"`
	CardCode    string         `json:"card_code" db:"card_code"`
	CardName    string         `json:"card_name" db:"card_name"`
	Set         SetModel       `json:"set" db:"set"`
	Supertype   SupertypeModel `json:"supertype" db:"supertype"`
	Types       []TypeModel    `json:"types" db:"types"`
	Subtypes    []SubtypeModel `json:"subtypes" db:"subtypes"`
	Rarity      RarityModel    `json:"rarity" db:"rarity"`
	MarketPrice float32        `json:"market_price" db:"market_price"`
	Pricetype   PricetypeModel `json:"pricetype" db:"pricetype"`
	Image       string         `json:"image" db:"image"`
}

type CardSearchParams struct {
	CardCode      string `json:"card_code" db:"card_code"`
	CardName      string `json:"card_name" db:"card_name"`
	SetName       string `json:"set_name" db:"set_name"`
	SupertypeName string `json:"supertype_name" db:"supertype_name"`
	RarityName    string `json:"rarity_name" db:"rarity_name"`
	PricetypeName string `json:"pricetype_name" db:"pricetype_name"`
}

type CardsList struct {
	TotalRecords int         `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"current_page"`
	Size         int         `json:"size"`
	Data         []CardModel `json:"data"`
}

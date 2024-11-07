package models

type CardEntity struct {
	CardID          int     `json:"card_id"`
	CardCode        string  `json:"card_code"`
	CardName        string  `json:"card_name"`
	SetID           int     `json:"set_id"`
	SupertypeID     int     `json:"supertype_id"`
	RarityID        int     `json:"rarity_id"`
	MarketPrice     float32 `json:"market_price"`
	PricetypeID     int     `json:"pricetype_id"`
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

type CardsList struct {
	TotalRecords int          `json:"total_records"`
	TotalPages   int          `json:"total_pages"`
	CurrentPage  int          `json:"current_page"`
	Size         int          `json:"size"`
	Data         []*CardModel `json:"data"`
}

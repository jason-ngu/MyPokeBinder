package models

type CardSubtypeEntity struct {
	CardID    int `json:"card_id" db:"card_id"`
	SubtypeID int `json:"subtype_id" db:"subtype_id"`
}

type CardSubtypeModel struct {
	CardID    int `json:"card_id" db:"card_id"`
	SubtypeID int `json:"subtype_id" db:"subtype_id"`
}

type CardSubtypesList struct {
	TotalRecords int                `json:"total_records"`
	TotalPages   int                `json:"total_pages"`
	CurrentPage  int                `json:"current_page"`
	Size         int                `json:"size"`
	Data         []CardSubtypeModel `json:"data"`
}

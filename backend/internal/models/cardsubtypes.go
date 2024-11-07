package models

type CardSubtypeEntity struct {
	CardID    int `json:"card_id"`
	SubtypeID int `json:"subtype_id"`
}

type CardSubtypeModel struct {
	CardID    int `json:"card_id"`
	SubtypeID int `json:"subtype_id"`
}

type CardSubtypesList struct {
	TotalRecords int                 `json:"total_records"`
	TotalPages   int                 `json:"total_pages"`
	CurrentPage  int                 `json:"current_page"`
	Size         int                 `json:"size"`
	Data         []*CardSubtypeModel `json:"data"`
}

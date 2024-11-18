package models

type CardTypeEntity struct {
	CardID int `json:"card_id" db:"card_id"`
	TypeID int `json:"type_id" db:"type_id"`
}

type CardTypeModel struct {
	CardID int `json:"card_id" db:"card_id"`
	TypeID int `json:"type_id" db:"type_id"`
}

type CardTypesList struct {
	TotalRecords int             `json:"total_records"`
	TotalPages   int             `json:"total_pages"`
	CurrentPage  int             `json:"current_page"`
	Size         int             `json:"size"`
	Data         []CardTypeModel `json:"data"`
}

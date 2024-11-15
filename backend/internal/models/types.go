package models

type TypeEntity struct {
	TypeID   int    `json:"type_id" db:"type_id"`
	TypeName string `json:"type_name" db:"type_name"`
}

type TypeModel struct {
	TypeID   int    `json:"type_id" db:"type_id"`
	TypeName string `json:"type_name" db:"type_name"`
}

type TypeSearchParams struct {
	TypeName string `json:"type_name" db:"type_name"`
}

type TypesList struct {
	TotalRecords int         `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"current_page"`
	Size         int         `json:"size"`
	Data         []TypeModel `json:"data"`
}

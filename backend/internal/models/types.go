package models

type TypeEntity struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

type TypeModel struct {
	TypeID   int
	TypeName string
}

type TypesList struct {
	TotalRecords int          `json:"total_records"`
	TotalPages   int          `json:"total_pages"`
	CurrentPage  int          `json:"current_page"`
	Size         int          `json:"size"`
	Types        []*TypeModel `json:"types"`
}

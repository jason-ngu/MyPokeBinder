package models

type SubtypeEntity struct {
	SubtypeID   int    `json:"subtype_id" db:"subtype_id"`
	SubtypeName string `json:"subtype_name" db:"subtype_name"`
}

type SubtypeModel struct {
	SubtypeID   int    `json:"subtype_id" db:"subtype_id"`
	SubtypeName string `json:"subtype_name" db:"subtype_name"`
}

type SubtypeSearchParams struct {
	SubtypeName string `json:"subtype_name" db:"subtype_name"`
}

type SubtypesList struct {
	TotalRecords int            `json:"total_records"`
	TotalPages   int            `json:"total_pages"`
	CurrentPage  int            `json:"current_page"`
	Size         int            `json:"size"`
	Data         []SubtypeModel `json:"data"`
}

package models

type SupertypeEntity struct {
	SupertypeID   int    `json:"supertype_id" db:"supertype_id"`
	SupertypeName string `json:"supertype_name" db:"supertype_name"`
}

type SupertypeModel struct {
	SupertypeID   int    `json:"supertype_id" db:"supertype_id"`
	SupertypeName string `json:"supertype_name" db:"supertype_name"`
}

type SupertypeSearchParams struct {
	SupertypeName string `json:"supertype_name" db:"supertype_name"`
}

type SupertypesList struct {
	TotalRecords int              `json:"total_records"`
	TotalPages   int              `json:"total_pages"`
	CurrentPage  int              `json:"current_page"`
	Size         int              `json:"size"`
	Data         []SupertypeModel `json:"data"`
}

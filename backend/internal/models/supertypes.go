package models

type SupertypeEntity struct {
	SupertypeID   int    `json:"supertype_id"`
	SupertypeName string `json:"supertype_name"`
}

type SupertypeModel struct {
	SupertypeID   int
	SupertypeName string
}

type SupertypesList struct {
	TotalRecords int               `json:"total_records"`
	TotalPages   int               `json:"total_pages"`
	CurrentPage  int               `json:"current_page"`
	Size         int               `json:"size"`
	Data         []*SupertypeModel `json:"data"`
}

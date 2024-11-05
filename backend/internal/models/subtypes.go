package models

type SubtypeEntity struct {
	SubtypeID   int    `json:"subtype_id"`
	SubtypeName string `json:"subtype_name"`
}

type SubtypeModel struct {
	SubtypeID   int
	SubtypeName string
}

type SubtypesList struct {
	TotalRecords int             `json:"total_records"`
	TotalPages   int             `json:"total_pages"`
	CurrentPage  int             `json:"current_page"`
	Size         int             `json:"size"`
	Data         []*SubtypeModel `json:"data"`
}

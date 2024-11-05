package models

type PricetypeEntity struct {
	PricetypeID   int    `json:"pricetype_id"`
	PricetypeName string `json:"pricetype_name"`
}

type PricetypeModel struct {
	PricetypeID   int
	PricetypeName string
}

type PricetypesList struct {
	TotalRecords int               `json:"total_records"`
	TotalPages   int               `json:"total_pages"`
	CurrentPage  int               `json:"current_page"`
	Size         int               `json:"size"`
	Data         []*PricetypeModel `json:"data"`
}

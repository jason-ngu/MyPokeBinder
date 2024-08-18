package models

type PricetypeEntity struct {
	PricetypeID   int    `json:"pricetype_id"`
	PricetypeName string `json:"pricetype_name"`
}

type PricetypeModel struct {
	PricetypeID   int
	PricetypeName string
}

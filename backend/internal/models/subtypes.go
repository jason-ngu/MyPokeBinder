package models

type SubtypeEntity struct {
	SubtypeID   int    `json:"subtype_id"`
	SubtypeName string `json:"subtype_name"`
}

type SubtypeModel struct {
	SubtypeID   int
	SubtypeName string
}

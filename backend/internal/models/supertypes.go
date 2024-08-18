package models

type SupertypeEntity struct {
	SupertypeID   int    `json:"supertype_id"`
	SupertypeName string `json:"supertype_name"`
}

type SupertypeModel struct {
	SupertypeID   int
	SupertypeName string
}

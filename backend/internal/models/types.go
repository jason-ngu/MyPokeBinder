package models

type TypeEntity struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

type TypeModel struct {
	TypeName string
}

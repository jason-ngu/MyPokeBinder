package models

type TypesEntity struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

type TypesModel struct {
	TypeName string
}

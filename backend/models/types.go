package models

type Types struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

type TypesRepository interface {
	SearchTypes() *Types
}

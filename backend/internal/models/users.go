package models

type UserEntity struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

type UserModel struct {
	UserName string
}

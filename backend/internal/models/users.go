package models

type UserEntity struct {
	UserID       int    `json:"user_id" db:"user_id"`
	Name         string `json:"name" db:"name"`
	ProviderType string `json:"provider_type" db:"provider_type"`
	ProviderKey  string `json:"provider_key" db:"provider_key"`
}

type UserModel struct {
	UserID       int    `json:"user_id" db:"user_id"`
	Name         string `json:"name" db:"name"`
	ProviderType string `json:"provider_type" db:"provider_type"`
	ProviderKey  string `json:"provider_key" db:"provider_key"`
}

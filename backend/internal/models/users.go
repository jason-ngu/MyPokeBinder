package models

type UserEntity struct {
	UserID       int    `json:"user_id"`
	Name         string `json:"name"`
	ProviderKey  string `json:"provider_key"`
	ProviderType string `json:"provider_type"`
}

type UserModel struct {
	UserID       int
	Name         string
	ProviderKey  string
	ProviderType string
}

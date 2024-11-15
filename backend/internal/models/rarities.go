package models

type RarityEntity struct {
	RarityID   int    `json:"rarity_id" db:"rarity_id"`
	RarityName string `json:"rarity_name" db:"rarity_name"`
}

type RarityModel struct {
	RarityID   int    `json:"rarity_id" db:"rarity_id"`
	RarityName string `json:"rarity_name" db:"rarity_name"`
}

type RaritySearchParams struct {
	RarityName string `json:"rarity_name" db:"rarity_name"`
}

type RaritiesList struct {
	TotalRecords int           `json:"total_records"`
	TotalPages   int           `json:"total_pages"`
	CurrentPage  int           `json:"current_page"`
	Size         int           `json:"size"`
	Data         []RarityModel `json:"data"`
}

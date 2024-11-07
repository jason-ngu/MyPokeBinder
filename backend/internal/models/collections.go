package models

type CollectionEntity struct {
	CollectionID   int    `json:"collection_id"`
	CollectionName string `json:"collection_name"`
	UserID         int    `json:"user_id"`
}

type CollectionModel struct {
	CollectionID   int
	CollectionName string
	UserID         int
}

type CollectionsList struct {
	TotalRecords int                `json:"total_records"`
	TotalPages   int                `json:"total_pages"`
	CurrentPage  int                `json:"current_page"`
	Size         int                `json:"size"`
	Data         []*CollectionModel `json:"data"`
}

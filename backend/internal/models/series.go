package models

type SeriesEntity struct {
	SeriesID   int    `json:"series_id"`
	SeriesName string `json:"series_name"`
}

type SeriesModel struct {
	SeriesID   int
	SeriesName string
}

type SeriesList struct {
	TotalRecords int            `json:"total_records"`
	TotalPages   int            `json:"total_pages"`
	CurrentPage  int            `json:"current_page"`
	Size         int            `json:"size"`
	Data         []*SeriesModel `json:"data"`
}

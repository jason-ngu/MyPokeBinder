package models

type SeriesEntity struct {
	SeriesID   int    `json:"series_id" db:"series_id"`
	SeriesName string `json:"series_name" db:"series_name"`
}

type SeriesModel struct {
	SeriesID   int    `json:"series_id" db:"series_id"`
	SeriesName string `json:"series_name" db:"series_name"`
}

type SeriesSearchParams struct {
	SeriesName string
}

type SeriesList struct {
	TotalRecords int           `json:"total_records"`
	TotalPages   int           `json:"total_pages"`
	CurrentPage  int           `json:"current_page"`
	Size         int           `json:"size"`
	Data         []SeriesModel `json:"data"`
}

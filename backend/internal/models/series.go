package models

type SeriesEntity struct {
	SeriesId   int    `json:"series_id"`
	SeriesName string `json:"series_name"`
}

type SeriesModel struct {
	SeriesName string
}

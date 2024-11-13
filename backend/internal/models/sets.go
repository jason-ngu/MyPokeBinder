package models

import "time"

type SetEntity struct {
	SetID             int       `json:"set_id"`
	SetCode           string    `json:"set_code"`
	SetName           string    `json:"set_name"`
	SeriesID          int       `json:"series_id"`
	PtcgoCode         string    `json:"ptcgo_code"`
	CardTotal         int       `json:"card_total"`
	ExtendedCardTotal int       `json:"extended_card_total"`
	SetReleaseDate    time.Time `json:"set_release_date"`
	SymbolImage       string    `json:"symbol_image"`
	LogoImage         string    `json:"logo_image"`
	SyncDateCreated   time.Time `json:"sync_date_created"`
	SyncDateUpdated   time.Time `json:"sync_date_updated"`
}

type SetModel struct {
	SetID             int
	SetName           string
	SetCode           string
	Series            *SeriesModel
	PtcgoCode         string
	CardTotal         int
	ExtendedCardTotal int
	SetReleaseDate    time.Time
	SymbolImage       string
	LogoImage         string
}

type SetSearchParams struct {
	SetName    string
	SetCode    string
	SeriesName string
	PtcgoCode  string
}

type SetsList struct {
	TotalRecords int         `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	CurrentPage  int         `json:"current_page"`
	Size         int         `json:"size"`
	Data         []*SetModel `json:"data"`
}

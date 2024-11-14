package models

import "time"

type SetEntity struct {
	SetID             int       `json:"set_id" db:"set_id"`
	SetCode           string    `json:"set_code" db:"set_code"`
	SetName           string    `json:"set_name" db:"set_name"`
	SeriesID          int       `json:"series_id" db:"series_id"`
	PtcgoCode         string    `json:"ptcgo_code" db:"ptcgo_code"`
	CardTotal         int       `json:"card_total" db:"card_total"`
	ExtendedCardTotal int       `json:"extended_card_total" db:"extended_card_total"`
	SetReleaseDate    time.Time `json:"set_release_date" db:"set_release_date"`
	SymbolImage       string    `json:"symbol_image" db:"symbol_image"`
	LogoImage         string    `json:"logo_image" db:"logo_image"`
	SyncDateCreated   time.Time `json:"sync_date_created" db:"sync_date_created"`
	SyncDateUpdated   time.Time `json:"sync_date_updated" db:"sync_date_updated"`
}

type SetModel struct {
	SetID             int         `json:"set_id" db:"set_id"`
	SetCode           string      `json:"set_code" db:"set_code"`
	SetName           string      `json:"set_name" db:"set_name"`
	Series            SeriesModel `json:"series" db:"series"`
	PtcgoCode         string      `json:"ptcgo_code" db:"ptcgo_code"`
	CardTotal         int         `json:"card_total" db:"card_total"`
	ExtendedCardTotal int         `json:"extended_card_total" db:"extended_card_total"`
	SetReleaseDate    time.Time   `json:"set_release_date" db:"set_release_date"`
	SymbolImage       string      `json:"symbol_image" db:"symbol_image"`
	LogoImage         string      `json:"logo_image" db:"logo_image"`
}

type SetSearchParams struct {
	SetName    string
	SetCode    string
	SeriesName string
	PtcgoCode  string
}

type SetsList struct {
	TotalRecords int        `json:"total_records"`
	TotalPages   int        `json:"total_pages"`
	CurrentPage  int        `json:"current_page"`
	Size         int        `json:"size"`
	Data         []SetModel `json:"data"`
}

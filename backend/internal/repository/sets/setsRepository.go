package setsRepository

import (
	"backend/internal/models"
	seriesRepository "backend/internal/repository/series"
	"database/sql"
	"time"
)

func GetAllSets(db *sql.DB) ([]models.SetModel, error) {
	rows, err := db.Query("SELECT * FROM public.sets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sets []models.SetModel

	for rows.Next() {
		var t models.SetEntity

		err := rows.Scan(&t.SetID, &t.SetCode, &t.SetName, &t.SeriesId, &t.PtcgoCode,
			&t.CardTotal, &t.ExtendedCardTotal, &t.SetReleaseDate, &t.SymbolImage, &t.LogoImage,
			&t.SyncDateCreated, &t.SyncDateUpdated)
		if err != nil {
			return nil, err
		}

		series, err := seriesRepository.GetSeriesById(db, t.SeriesId)
		if err != nil {
			return nil, err
		}

		sets = append(sets, models.SetModel{
			SetName:           t.SetName,
			SetCode:           t.SetCode,
			SeriesName:        series.SeriesName,
			PtcgoCode:         t.PtcgoCode,
			CardTotal:         t.CardTotal,
			ExtendedCardTotal: t.ExtendedCardTotal,
			SetReleaseDate:    t.SetReleaseDate,
			SymbolImage:       t.SymbolImage,
			LogoImage:         t.LogoImage,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sets, nil
}

func GetSetById(db *sql.DB, id int) (models.SetModel, error) {
	row := db.QueryRow("SELECT * FROM public.sets WHERE set_id = $1", id)
	if err := row.Err(); err != nil {
		return models.SetModel{}, err
	}

	var t models.SetEntity
	err := row.Scan(&t.SetID, &t.SetCode, &t.SetName, &t.SeriesId, &t.PtcgoCode,
		&t.CardTotal, &t.ExtendedCardTotal, &t.SetReleaseDate, &t.SymbolImage, &t.LogoImage,
		&t.SyncDateCreated, &t.SyncDateUpdated)
	if err != nil {
		return models.SetModel{}, err
	}

	series, err := seriesRepository.GetSeriesById(db, t.SeriesId)
	if err != nil {
		return models.SetModel{}, err
	}

	return models.SetModel{
		SetName:           t.SetName,
		SetCode:           t.SetCode,
		SeriesName:        series.SeriesName,
		PtcgoCode:         t.PtcgoCode,
		CardTotal:         t.CardTotal,
		ExtendedCardTotal: t.ExtendedCardTotal,
		SetReleaseDate:    t.SetReleaseDate,
		SymbolImage:       t.SymbolImage,
		LogoImage:         t.LogoImage,
	}, nil
}

func GetSetByName(db *sql.DB, setName string) (models.SetModel, error) {
	row := db.QueryRow("SELECT * FROM public.sets WHERE set_name = $1", setName)
	if err := row.Err(); err != nil {
		return models.SetModel{}, err
	}

	var t models.SetEntity
	err := row.Scan(&t.SetID, &t.SetCode, &t.SetName, &t.SeriesId, &t.PtcgoCode,
		&t.CardTotal, &t.ExtendedCardTotal, &t.SetReleaseDate, &t.SymbolImage, &t.LogoImage,
		&t.SyncDateCreated, &t.SyncDateUpdated)
	if err != nil {
		return models.SetModel{}, err
	}

	series, err := seriesRepository.GetSeriesById(db, t.SeriesId)
	if err != nil {
		return models.SetModel{}, err
	}

	return models.SetModel{
		SetName:           t.SetName,
		SetCode:           t.SetCode,
		SeriesName:        series.SeriesName,
		PtcgoCode:         t.PtcgoCode,
		CardTotal:         t.CardTotal,
		ExtendedCardTotal: t.ExtendedCardTotal,
		SetReleaseDate:    t.SetReleaseDate,
		SymbolImage:       t.SymbolImage,
		LogoImage:         t.LogoImage,
	}, nil
}

func CreateSet(db *sql.DB, newSet models.SetModel) (models.SetModel, error) {
	set, err := GetSetByName(db, newSet.SetName)
	if (err == nil) && (set != models.SetModel{}) {
		return set, err
	}

	syncDateCreated := time.Now()
	syncDateUpdated := time.Now()

	series, err := seriesRepository.GetSeriesByName(db, newSet.SeriesName)
	if err != nil {
		return models.SetModel{}, err
	}

	_, err = db.Exec(`INSERT INTO public.sets 
		(set_code, set_name, series_id, ptcgo_code, card_total, extended_card_total, set_release_date, symbol_image, logo_image, sync_date_created, sync_date_updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		newSet.SetCode, newSet.SetName, series.SeriesID, newSet.PtcgoCode, newSet.CardTotal, newSet.ExtendedCardTotal, newSet.SetReleaseDate, newSet.SymbolImage, newSet.LogoImage, syncDateCreated, syncDateUpdated)
	if err != nil {
		return models.SetModel{}, err
	}

	createdSet, err := GetSetByName(db, newSet.SetName)
	if err != nil {
		return models.SetModel{}, err
	}

	return createdSet, nil
}

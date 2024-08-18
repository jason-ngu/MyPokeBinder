package seriesRepository

import (
	"backend/internal/models"
	"database/sql"
)

func GetAllSeries(db *sql.DB) ([]models.SeriesModel, error) {
	rows, err := db.Query("SELECT * FROM public.series")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var series []models.SeriesModel

	for rows.Next() {
		var t models.SeriesEntity

		err := rows.Scan(&t.SeriesID, &t.SeriesName)
		if err != nil {
			return nil, err
		}

		series = append(series, models.SeriesModel(t))
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return series, nil
}

func GetSeriesById(db *sql.DB, id int) (models.SeriesModel, error) {
	row := db.QueryRow("SELECT * FROM public.series WHERE series_id = $1", id)

	var t models.SeriesEntity
	err := row.Scan(&t.SeriesID, &t.SeriesName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.SeriesModel{}, nil
		}
		return models.SeriesModel{}, err
	}

	return models.SeriesModel(t), nil
}

func GetSeriesByName(db *sql.DB, seriesName string) (models.SeriesModel, error) {
	row := db.QueryRow("SELECT * FROM public.series WHERE series_name = $1", seriesName)

	var t models.SeriesEntity
	err := row.Scan(&t.SeriesID, &t.SeriesName)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.SeriesModel{}, nil
		}
		return models.SeriesModel{}, err
	}

	return models.SeriesModel(t), nil
}

func CreateSeries(db *sql.DB, newSeries models.SeriesModel) (models.SeriesModel, error) {
	series, err := GetSeriesByName(db, newSeries.SeriesName)
	// If series already exists, return it
	if (err == nil) && (series != models.SeriesModel{}) {
		return series, err
	}

	_, err = db.Exec("INSERT INTO public.series (series_name) VALUES ($1)", newSeries.SeriesName)
	if err != nil {
		return models.SeriesModel{}, err
	}

	createdSeries, err := GetSeriesByName(db, newSeries.SeriesName)
	if err != nil {
		return models.SeriesModel{}, err
	}

	return createdSeries, nil
}

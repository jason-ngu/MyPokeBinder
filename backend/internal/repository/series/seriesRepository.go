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

		series = append(series, models.SeriesModel{SeriesName: t.SeriesName})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return series, nil
}

func GetSeriesById(db *sql.DB, id int) (models.SeriesModel, error) {
	row := db.QueryRow("SELECT * FROM public.series WHERE series_id = $1", id)
	if err := row.Err(); err != nil {
		return models.SeriesModel{}, err
	}

	var t models.SeriesEntity
	err := row.Scan(&t.SeriesID, &t.SeriesName)
	if err != nil {
		return models.SeriesModel{}, err
	}

	return models.SeriesModel{SeriesName: t.SeriesName}, nil
}

func GetSeriesByName(db *sql.DB, seriesName string) (models.SeriesModel, error) {
	row := db.QueryRow("SELECT * FROM public.series WHERE series_name = $1", seriesName)
	if err := row.Err(); err != nil {
		return models.SeriesModel{}, err
	}

	var t models.SeriesEntity
	err := row.Scan(&t.SeriesID, &t.SeriesName)
	if err != nil {
		return models.SeriesModel{}, err
	}

	return models.SeriesModel{SeriesName: t.SeriesName}, nil
}

func CreateSeries(db *sql.DB, newSeries models.SeriesModel) (int64, error) {
	series, err := GetSeriesByName(db, newSeries.SeriesName)
	if (series != models.SeriesModel{}) && (err != nil) {
		return 0, err
	}

	result, err := db.Exec("INSERT INTO public.series (series_name) VALUES ($1)", newSeries.SeriesName)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

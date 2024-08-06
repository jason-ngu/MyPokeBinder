package seriesService

import (
	internal "backend/internal"
	"backend/internal/models"
	seriesRepository "backend/internal/repository/series"
)

func GetAllSeries(env *internal.Env) ([]models.SeriesModel, error) {
	allSeries, err := seriesRepository.GetAllSeries(env.DB)
	if err != nil {
		return []models.SeriesModel{}, err
	}
	return allSeries, nil
}

func GetSeriesById(env *internal.Env, id int) (models.SeriesModel, error) {
	series, err := seriesRepository.GetSeriesById(env.DB, id)
	if err != nil {
		return models.SeriesModel{}, err
	}
	return series, nil
}

func GetSeriesByName(env *internal.Env, seriesName string) (models.SeriesModel, error) {
	series, err := seriesRepository.GetSeriesByName(env.DB, seriesName)
	if err != nil {
		return models.SeriesModel{}, err
	}
	return series, nil
}

func CreateSeries(env *internal.Env, newSeries models.SeriesModel) (models.SeriesModel, error) {
	series, err := seriesRepository.CreateSeries(env.DB, newSeries)
	if err != nil {
		return models.SeriesModel{}, err
	}
	return series, nil
}

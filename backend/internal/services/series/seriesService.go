package seriesService

import (
	internal "backend/internal"
	"backend/internal/models"
	seriesRepository "backend/internal/repository/series"
)

func GetAllSeries(env *internal.Env) ([]models.SeriesModel, error) {
	allTypes, err := seriesRepository.GetAllSeries(env.DB)
	if err != nil {
		return []models.SeriesModel{}, err
	}
	return allTypes, nil
}

func GetSeriesById(env *internal.Env, id int) (models.SeriesModel, error) {
	t, err := seriesRepository.GetSeriesById(env.DB, id)
	if err != nil {
		return models.SeriesModel{}, err
	}
	return t, nil
}

func CreateSeries(env *internal.Env, newSeries models.SeriesModel) (int64, error) {
	id, err := seriesRepository.CreateSeries(env.DB, newSeries)
	if err != nil {
		return 0, err
	}
	return id, nil
}

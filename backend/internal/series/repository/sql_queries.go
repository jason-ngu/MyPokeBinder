package repository

const (
	createSeries           = `INSERT INTO public.series (series_name) VALUES ($1) RETURN *`
	getSeriesById          = `SELECT * FROM public.series WHERE series_id = $1`
	getAllSeries           = `SELECT * FROM public.series`
	getTotalCountAllSeries = `SELECT COUNT(*) FROM public.series`
)

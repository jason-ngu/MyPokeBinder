package repository

const (
	createSeries           = `INSERT INTO public.series (series_name) VALUES ($1) RETURNING *`
	getSeriesById          = `SELECT * FROM public.series WHERE series_id = $1`
	getAllSeries           = `SELECT * FROM public.series%s OFFSET %d LIMIT %d`
	getTotalCountAllSeries = `SELECT COUNT(*) FROM public.series%s`
)

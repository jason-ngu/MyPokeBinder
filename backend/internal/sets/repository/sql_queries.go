package repository

const (
	createSet = `INSERT INTO public.sets 
		(set_code, set_name, series_id, ptcgo_code, card_total, extended_card_total, set_release_date, symbol_image, logo_image, sync_date_created, sync_date_updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING *`
	getSetById = `SELECT * FROM public.sets WHERE set_id = $1`
	getAllSets = `SELECT * FROM public.sets s
							JOIN public.series ss
							ON s.series_id = ss.series_id`
	getTotalCountAllSets = `SELECT COUNT(*) FROM public.sets`
	deleteSet            = `DELETE FROM public.sets WHERE set_id = $1`
)

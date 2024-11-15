package repository

const (
	createSet = `INSERT INTO public.sets 
		(set_code, set_name, series_id, ptcgo_code, card_total, extended_card_total, set_release_date, symbol_image, logo_image, sync_date_created, sync_date_updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING set_id`
	getSetById = `SELECT s.set_id,
							s.set_code,
							s.set_name,
							s.ptcgo_code,
							s.card_total,
							s.extended_card_total,
							s.set_release_date,
							s.symbol_image,
							s.logo_image,
							sr.series_id AS "series.series_id",
							sr.series_name AS "series.series_name"
							FROM public.sets s
							JOIN public.series sr
							ON s.series_id = sr.series_id
							WHERE set_id = $1`
	getAllSets = `SELECT 	s.set_id,
							s.set_code,
							s.set_name,
							s.ptcgo_code,
							s.card_total,
							s.extended_card_total,
							s.set_release_date,
							s.symbol_image,
							s.logo_image,
							sr.series_id AS "series.series_id",
							sr.series_name AS "series.series_name"
							FROM public.sets s
							JOIN public.series sr
							ON s.series_id = sr.series_id
							WHERE 	s.set_name = :set_name AND
									s.set_code = :set_code AND
									sr.series_name = :series_name AND
									s.ptcgo_code = :ptcgo_code
							OFFSET %d LIMIT %d`
	getTotalCountAllSets = `SELECT COUNT(*) FROM public.sets s
							JOIN public.series sr
							ON s.series_id = sr.series_id
							WHERE 	s.set_name = :set_name AND
									s.set_code = :set_code AND
									sr.series_name = :series_name AND
									s.ptcgo_code = :ptcgo_code`
	deleteSet = `DELETE FROM public.sets WHERE set_id = $1`
)

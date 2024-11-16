package repository

const (
	createPricetype            = `INSERT INTO public.pricetypes (pricetype_name) VALUES ($1) RETURNING pricetype_id`
	getPricetypeById           = `SELECT * FROM public.pricetypes WHERE pricetype_id = $1`
	getAllPricetypes           = `SELECT * FROM public.pricetypes WHERE pricetype_name = :pricetype_name OFFSET %d LIMIT %d`
	getTotalCountAllPricetypes = `SELECT COUNT(*) FROM public.pricetypes WHERE pricetype_name = :pricetype_name`
)

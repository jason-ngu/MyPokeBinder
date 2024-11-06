package repository

const (
	createPricetype            = `INSERT INTO public.pricetypes (pricetype_name) VALUES ($1) RETURNING *`
	getPricetypeById           = `SELECT * FROM public.pricetypes WHERE pricetype_id = $1`
	getAllPricetypes           = `SELECT * FROM public.pricetypes`
	getTotalCountAllPricetypes = `SELECT COUNT(*) FROM public.pricetypes`
)

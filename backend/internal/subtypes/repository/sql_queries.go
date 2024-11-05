package repository

const (
	createSubtypeQuery       = `INSERT INTO public.subtypes (subtype_name) VALUES ($1) RETURN *`
	getSubtypeById           = `SELECT * FROM public.subtypes WHERE subtype_id = $1`
	getAllSubtypes           = `SELECT * FROM public.subtypes`
	getTotalCountAllSubtypes = `SELECT COUNT(*) FROM public.subtypes`
)

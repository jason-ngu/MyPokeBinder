package repository

const (
	createSubtype            = `INSERT INTO public.subtypes (subtype_name) VALUES ($1) RETURNING subtype_id`
	getSubtypeById           = `SELECT * FROM public.subtypes WHERE subtype_id = $1`
	getAllSubtypes           = `SELECT * FROM public.subtypes WHERE subtype_name = :subtype_name OFFSET %d LIMIT %d`
	getTotalCountAllSubtypes = `SELECT COUNT(*) FROM public.subtypes WHERE subtype_name = :subtype_name`
)

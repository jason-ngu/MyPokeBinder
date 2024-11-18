package repository

const (
	createSupertype            = `INSERT INTO public.supertypes (supertype_name) VALUES ($1) RETURNING supertype_id`
	getSupertypeById           = `SELECT * FROM public.supertypes WHERE supertype_id = $1`
	getAllSupertypes           = `SELECT * FROM public.supertypes`
	getTotalCountAllSupertypes = `SELECT COUNT(*) FROM public.supertypes`
)

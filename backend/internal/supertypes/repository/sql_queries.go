package repository

const (
	createSupertype            = `INSERT INTO public.supertypes (supertype_name) VALUES ($1) RETURN *`
	getSupertypeById           = `SELECT * FROM public.supertypes WHERE supertype_id = $1`
	getAllSupertypes           = `SELECT * FROM public.supertypes`
	getTotalCountAllSupertypes = `SELECT COUNT(*) FROM public.supertypes`
)

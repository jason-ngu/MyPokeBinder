package repository

const (
	createTypeQuery       = `INSERT INTO public.types (type_name) VALUES ($1) RETURN *`
	getTypeById           = `SELECT * FROM public.types WHERE type_id = $1`
	getAllTypes           = `SELECT * FROM public.types`
	getTotalCountAllTypes = `SELECT COUNT(type_id) FROM public.types`
)

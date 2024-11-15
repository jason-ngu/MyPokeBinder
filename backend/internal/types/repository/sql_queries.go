package repository

const (
	createType            = `INSERT INTO public.types (type_name) VALUES ($1) RETURNING *`
	getTypeById           = `SELECT * FROM public.types WHERE type_id = $1`
	getAllTypes           = `SELECT * FROM public.types WHERE type_name = :type_name OFFSET %d LIMIT %d`
	getTotalCountAllTypes = `SELECT COUNT(*) FROM public.types WHERE type_name = :type_name`
)

package repository

const (
	createCollection = `INSERT INTO public.collections
						(collection_name, user_id)
						VALUES ($1, $2)`
	getCollectionById      = `SELECT * FROM public.collections WHERE collection_id = $1`
	getCollectionsByUserID = `SELECT * FROM public.collections WHERE user_id = $1`
	getTotalCountByUserID  = `SELECT COUNT(*) FROM public.collections WHERE user_id = $1`
	updateCollection       = `UPDATE public.collections
								SET collection_name = $1
								WHERE collection_id = $2 AND user_id = $3
								RETURNING collection_id`
	deleteCollection = `DELETE FROM public.collections WHERE collection_id = $1`
)

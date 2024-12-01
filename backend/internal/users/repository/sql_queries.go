package repository

const (
	createUser = `INSERT into public.users
					(name, provider_key, provider_type)
					VALUES ($1, $2, $3)
					RETURNING user_id`
	getUserById       = `SELECT * FROM public.users WHERE user_id = $1`
	getUserByProvider = `SELECT * FROM public.users WHERE provider_type = $1 AND provider_key = $2`
)

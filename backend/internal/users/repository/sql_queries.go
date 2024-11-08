package repository

const (
	createUser = `INSERT into public.users
					(name, provider_key, provider_type)
					VALUES ($1, $2, $3)`
	getUserById = `SELECT * FROM public.users WHERE user_id = $1`
)

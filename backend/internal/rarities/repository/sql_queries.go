package repository

const (
	createRarity             = `INSERT INTO public.rarities (rarity_name) VALUES ($1) RETURNING *`
	getRarityById            = `SELECT * FROM public.rarities WHERE rarity_id = $1`
	getAllRarities           = `SELECT * FROM public.rarities`
	getTotalCountAllRarities = `SELECT COUNT(*) FROM public.rarities`
)

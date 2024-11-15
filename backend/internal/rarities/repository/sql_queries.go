package repository

const (
	createRarity             = `INSERT INTO public.rarities (rarity_name) VALUES ($1) RETURNING rarity_id`
	getRarityById            = `SELECT * FROM public.rarities WHERE rarity_id = $1`
	getAllRarities           = `SELECT * FROM public.rarities WHERE rarity_name = :rarity_name OFFSET %d LIMIT %d`
	getTotalCountAllRarities = `SELECT COUNT(*) FROM public.rarities WHERE rarity_name = :rarity_name`
)

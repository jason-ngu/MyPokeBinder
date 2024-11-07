package repository

const (
	createCardType        = `INSERT INTO public.cardtypes (card_id, type_id) VALUES ($1, $2) RETURNING *`
	getCardTypeByCardID   = `SELECT * FROM public.cardtypes WHERE card_id = $1`
	getTotalCountByCardID = `SELECT COUNT(*) FROM public.cardtypes WHERE card_id = $1`
)

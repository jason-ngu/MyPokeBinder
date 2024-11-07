package repository

const (
	createCardSubtype      = `INSERT INTO public.cardsubtypes (card_id, subtype_id) VALUES ($1, $2) RETURNING *`
	getCardSubtypeByCardID = `SELECT * FROM public.cardsubtypes WHERE card_id = $1`
	getTotalCountByCardID  = `SELECT COUNT(*) FROM public.cardsubtypes WHERE card_id = $1`
)

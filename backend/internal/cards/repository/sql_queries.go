package repository

const (
	createCard = `INSERT INTO public.cards
		(card_code, card_name, set_id, supertype_id, rarity_id, market_price, pricetype_id, image, sync_date_created, sync_date_updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	getCardById = `SELECT * FROM public.cards WHERE card_id = $1`
	getAllCards = `SELECT c.* FROM public.cards c
					JOIN public.sets s ON c.set_id = s.set_id
					JOIN public.supertypes st ON c.supertype_id = st.supertype_id
					JOIN public.rarities r ON c.rarity_id = r.rarity_id
					JOIN public.pricetypes pt ON c.pricetype_id = pt.pricetype_id`
	getTotalCountAllCards = `SELECT COUNT(DISTINCT card_code || pricetype_id) FROM public.cards`
	updateCard            = `UPDATE public.cards
		SET market_price = $1, sync_date_updated = $2
		WHERE card_id = $3
		RETURNING *`
)

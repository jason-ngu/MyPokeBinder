package repository

const (
	createCard = `INSERT INTO public.cards
		(card_code, card_name, set_id, supertype_id, rarity_id, market_price, pricetype_id, image, sync_date_created, sync_date_updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING card_id`
	getCardById = `SELECT 	c.card_id,
							c.card_code,
							c.card_name,
							s.set_id as "set.set_id",
							s.set_name as "set.set_name",
							st.supertype_id as "supertype.supertype_id",
							st.supertype_name as "supertype.supertype_name",
							r.rarity_id as "rarity.rarity_id",
							r.rarity_name as "rarity.rarity_name",
							c.market_price,
							pt.pricetype_id as "pricetype.pricetype_id",
							pt.pricetype_name as "pricetype.pricetype_name",
							c.image
					FROM public.cards c
					JOIN public.sets s ON c.set_id = s.set_id
					JOIN public.supertypes st ON c.supertype_id = st.supertype_id
					JOIN public.rarities r ON c.rarity_id = r.rarity_id
					JOIN public.pricetypes pt ON c.pricetype_id = pt.pricetype_id
					WHERE card_id = $1`
	getAllCards = `SELECT 	c.card_id,
							c.card_code,
							c.card_name,
							s.set_id as "set.set_id",
							s.set_name as "set.set_name",
							st.supertype_id as "supertype.supertype_id",
							st.supertype_name as "supertype.supertype_name",
							r.rarity_id as "rarity.rarity_id",
							r.rarity_name as "rarity.rarity_name",
							c.market_price,
							pt.pricetype_id as "pricetype.pricetype_id",
							pt.pricetype_name as "pricetype.pricetype_name",
							c.image
					FROM public.cards c
					JOIN public.sets s ON c.set_id = s.set_id
					JOIN public.supertypes st ON c.supertype_id = st.supertype_id
					JOIN public.rarities r ON c.rarity_id = r.rarity_id
					JOIN public.pricetypes pt ON c.pricetype_id = pt.pricetype_id`
	getTotalCountAllCards = `SELECT COUNT(DISTINCT c.card_code || c.pricetype_id) FROM public.cards c
								JOIN public.sets s ON c.set_id = s.set_id
								JOIN public.supertypes st ON c.supertype_id = st.supertype_id
								JOIN public.rarities r ON c.rarity_id = r.rarity_id
								JOIN public.pricetypes pt ON c.pricetype_id = pt.pricetype_id`
	updateCard = `UPDATE public.cards
		SET market_price = $1, sync_date_updated = $2
		WHERE card_id = $3
		RETURNING card_id`

	createCardType   = `INSERT INTO public.cardtypes (card_id, type_id) VALUES ($1, $2) RETURNING *`
	getCardtypesById = `SELECT 	t.type_id
								t.type_name
						FROM public.cardtypes ct
						JOIN public.types t ON ct.type_id = t.type_id
						WHERE ct.card_id = $1`

	createCardSubtype   = `INSERT INTO public.cardsubtypes (card_id, subtype_id) VALUES ($1, $2) RETURNING *`
	getCardsubtypesById = `SELECT 	st.subtype_id
									st.subtype_name
						FROM public.cardsubtypes cst
						JOIN public.subtypes st ON cst.subtype_id = st.subtype_id
						WHERE cst.card_id = $1`
)

package cardsRepository

import (
	"backend/common"
	"backend/internal/models"
	pricetypesRepository "backend/internal/repository/pricetypes"
	raritiesRepository "backend/internal/repository/rarities"
	setsRepository "backend/internal/repository/sets"
	supertypesRepository "backend/internal/repository/supertypes"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

var defaultLimit = 100
var defaultOffset = 0

func GetAllCards(db *sql.DB) ([]models.CardModel, error) {
	rows, err := db.Query("SELECT * FROM public.cards LIMIT $1 OFFSET $2", defaultLimit, defaultOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []models.CardModel

	for rows.Next() {
		var t models.CardEntity

		err := rows.Scan(&t.CardID, &t.CardCode, &t.CardName, &t.SetId, &t.SupertypeId,
			&t.RarityId, &t.MarketPrice, &t.PricetypeId, &t.Image, &t.SyncDateCreated, &t.SyncDateUpdated)
		if err != nil {
			return nil, err
		}

		set, err := setsRepository.GetSetById(db, t.SetId)
		if err != nil {
			return nil, err
		}

		supertype, err := supertypesRepository.GetSupertypeById(db, t.SupertypeId)
		if err != nil {
			return nil, err
		}

		rarity, err := raritiesRepository.GetRarityById(db, t.RarityId)
		if err != nil {
			return nil, err
		}

		pricetype, err := pricetypesRepository.GetPricetypeById(db, t.PricetypeId)
		if err != nil {
			return nil, err
		}

		cards = append(cards, models.CardModel{
			CardID:        t.CardID,
			CardName:      t.CardName,
			CardCode:      t.CardCode,
			SetName:       set.SetName,
			SupertypeName: supertype.SupertypeName,
			RarityName:    rarity.RarityName,
			MarketPrice:   t.MarketPrice,
			PricetypeName: pricetype.PricetypeName,
			Image:         t.Image,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func SearchCards(db *sql.DB, searchParams models.CardSearchParams) ([]models.CardModel, error) {
	sqlQuery := `SELECT c.* FROM public.cards c
					JOIN public.sets s ON c.set_id = s.set_id
					JOIN public.supertypes st ON c.supertype_id = st.supertype_id
					JOIN public.rarities r ON c.rarity_id = r.rarity_id
					JOIN public.pricetypes pt ON c.pricetype_id = pt.pricetype_id`
	var whereFilters []string
	if (searchParams != models.CardSearchParams{}) {
		sqlQuery += " WHERE "
		if searchParams.CardName != "" {
			formattedCardName := common.FormatStringForDatabase(searchParams.CardName)
			whereFilters = append(whereFilters, fmt.Sprintf("card_name = '%s'", formattedCardName))
		}
		if searchParams.CardCode != "" {
			whereFilters = append(whereFilters, fmt.Sprintf("card_code = '%s'", searchParams.CardCode))
		}
		if searchParams.SetName != "" {
			whereFilters = append(whereFilters, fmt.Sprintf("set_name = '%s'", searchParams.SetName))
		}
		if searchParams.SupertypeName != "" {
			whereFilters = append(whereFilters, fmt.Sprintf("supertype_name = '%s'", searchParams.SupertypeName))
		}
		if searchParams.RarityName != "" {
			whereFilters = append(whereFilters, fmt.Sprintf("rarity_name = '%s'", searchParams.RarityName))
		}
		if searchParams.PricetypeName != "" {
			whereFilters = append(whereFilters, fmt.Sprintf("pricetype_name = '%s'", searchParams.PricetypeName))
		}
		sqlQuery += strings.Join(whereFilters, " AND ")
	}

	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []models.CardModel

	for rows.Next() {
		var t models.CardEntity

		err := rows.Scan(&t.CardID, &t.CardCode, &t.CardName, &t.SetId, &t.SupertypeId,
			&t.RarityId, &t.MarketPrice, &t.PricetypeId, &t.Image, &t.SyncDateCreated, &t.SyncDateUpdated)
		if err != nil {
			return nil, err
		}

		set, err := setsRepository.GetSetById(db, t.SetId)
		if err != nil {
			return nil, err
		}

		supertype, err := supertypesRepository.GetSupertypeById(db, t.SupertypeId)
		if err != nil {
			return nil, err
		}

		rarity, err := raritiesRepository.GetRarityById(db, t.RarityId)
		if err != nil {
			return nil, err
		}

		pricetype, err := pricetypesRepository.GetPricetypeById(db, t.PricetypeId)
		if err != nil {
			return nil, err
		}

		cards = append(cards, models.CardModel{
			CardID:        t.CardID,
			CardName:      t.CardName,
			CardCode:      t.CardCode,
			SetName:       set.SetName,
			SupertypeName: supertype.SupertypeName,
			RarityName:    rarity.RarityName,
			MarketPrice:   t.MarketPrice,
			PricetypeName: pricetype.PricetypeName,
			Image:         t.Image,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func GetCardsByCardCode(db *sql.DB, cardCode string) ([]models.CardModel, error) {
	rows, err := db.Query("SELECT * FROM public.cards WHERE card_code = $1 LIMIT $2 OFFSET $3", cardCode, defaultLimit, defaultOffset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []models.CardModel

	for rows.Next() {
		var t models.CardEntity

		err := rows.Scan(&t.CardID, &t.CardCode, &t.CardName, &t.SetId, &t.SupertypeId,
			&t.RarityId, &t.MarketPrice, &t.PricetypeId, &t.Image, &t.SyncDateCreated, &t.SyncDateUpdated)
		if err != nil {
			return nil, err
		}

		set, err := setsRepository.GetSetById(db, t.SetId)
		if err != nil {
			return nil, err
		}

		supertype, err := supertypesRepository.GetSupertypeById(db, t.SupertypeId)
		if err != nil {
			return nil, err
		}

		rarity, err := raritiesRepository.GetRarityById(db, t.RarityId)
		if err != nil {
			return nil, err
		}

		pricetype, err := pricetypesRepository.GetPricetypeById(db, t.PricetypeId)
		if err != nil {
			return nil, err
		}

		cards = append(cards, models.CardModel{
			CardID:        t.CardID,
			CardName:      t.CardName,
			CardCode:      t.CardCode,
			SetName:       set.SetName,
			SupertypeName: supertype.SupertypeName,
			RarityName:    rarity.RarityName,
			MarketPrice:   t.MarketPrice,
			PricetypeName: pricetype.PricetypeName,
			Image:         t.Image,
		})
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func GetCard(db *sql.DB, cardCode string, pricetype string) (models.CardModel, error) {
	row := db.QueryRow(`SELECT c.* 
						FROM public.cards c 
						JOIN public.pricetypes pt ON c.pricetype_id = pt.pricetype_id
						WHERE c.card_code = $1 AND pt.pricetype_name = $2`,
		cardCode, pricetype)

	var t models.CardEntity
	err := row.Scan(&t.CardID, &t.CardCode, &t.CardName, &t.SetId, &t.SupertypeId,
		&t.RarityId, &t.MarketPrice, &t.PricetypeId, &t.Image, &t.SyncDateCreated, &t.SyncDateUpdated)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.CardModel{}, nil
		}
		return models.CardModel{}, err
	}

	set, err := setsRepository.GetSetById(db, t.SetId)
	if err != nil {
		return models.CardModel{}, err
	}

	supertype, err := supertypesRepository.GetSupertypeById(db, t.SupertypeId)
	if err != nil {
		return models.CardModel{}, err
	}

	rarity, err := raritiesRepository.GetRarityById(db, t.RarityId)
	if err != nil {
		return models.CardModel{}, err
	}

	return models.CardModel{
		CardID:        t.CardID,
		CardName:      t.CardName,
		CardCode:      cardCode,
		SetName:       set.SetName,
		SupertypeName: supertype.SupertypeName,
		RarityName:    rarity.RarityName,
		MarketPrice:   t.MarketPrice,
		PricetypeName: pricetype,
		Image:         t.Image,
	}, nil
}

func CreateCard(db *sql.DB, newCard models.CardModel) (models.CardModel, error) {
	syncDateCreated := time.Now()
	syncDateUpdated := time.Now()

	set, err := setsRepository.GetSetByName(db, newCard.SetName)
	if err != nil {
		return models.CardModel{}, err
	}

	supertype, err := supertypesRepository.GetSupertypeByName(db, newCard.SupertypeName)
	if err != nil {
		return models.CardModel{}, err
	}

	// Hardcode rarity to be promo if it is missing
	if newCard.RarityName == "" {
		newCard.RarityName = "Promo"
	}
	rarity, err := raritiesRepository.GetRarityByName(db, newCard.RarityName)
	if err != nil {
		return models.CardModel{}, err
	}

	pricetype, err := pricetypesRepository.GetPricetypeByName(db, newCard.PricetypeName)
	if err != nil {
		return models.CardModel{}, err
	}

	_, err = db.Exec(`INSERT INTO public.cards
		(card_code, card_name, set_id, supertype_id, rarity_id, market_price, pricetype_id, image, sync_date_created, sync_date_updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		newCard.CardCode, newCard.CardName, set.SetID, supertype.SupertypeID, rarity.RarityID, newCard.MarketPrice, pricetype.PricetypeID, newCard.Image, syncDateCreated, syncDateUpdated)
	if err != nil {
		return models.CardModel{}, err
	}

	createdCard, err := SearchCards(db, models.CardSearchParams{
		CardName:      newCard.CardName,
		CardCode:      newCard.CardCode,
		PricetypeName: newCard.PricetypeName,
	})
	if err != nil {
		return models.CardModel{}, err
	}

	return createdCard[0], nil
}

func UpdateCard(db *sql.DB, cardToUpdate models.CardModel) (models.CardModel, error) {
	syncDateUpdated := time.Now()

	pricetype, err := pricetypesRepository.GetPricetypeByName(db, cardToUpdate.PricetypeName)
	if err != nil {
		return models.CardModel{}, err
	}

	_, err = db.Exec(`UPDATE public.cards
		SET market_price = $1, sync_date_updated = $2
		WHERE card_code = $3 AND pricetype_id = $4`,
		cardToUpdate.MarketPrice, syncDateUpdated, cardToUpdate.CardCode, pricetype.PricetypeID)
	if err != nil {
		return models.CardModel{}, err
	}

	updatedCard, err := SearchCards(db, models.CardSearchParams{
		CardName:      cardToUpdate.CardName,
		CardCode:      cardToUpdate.CardCode,
		PricetypeName: cardToUpdate.PricetypeName,
	})
	if err != nil {
		return models.CardModel{}, err
	}

	return updatedCard[0], nil
}

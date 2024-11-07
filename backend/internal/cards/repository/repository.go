package repository

import (
	"backend/common"
	"backend/internal/cards"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type cardsRepo struct {
	db *sql.DB
}

func NewCardsRepository(db *sql.DB) cards.Repository {
	return &cardsRepo{db: db}
}

func (r *cardsRepo) Create(ctx context.Context, newCard *models.CardEntity) (*models.CardEntity, error) {
	syncDateCreated := time.Now()
	syncDateUpdated := time.Now()

	card := &models.CardEntity{}
	row := r.db.QueryRowContext(ctx, createCard, &newCard.CardCode, &newCard.CardName, &newCard.SetID, &newCard.SupertypeID, &newCard.RarityID, &newCard.MarketPrice, &newCard.PricetypeID, &newCard.Image, syncDateCreated, syncDateUpdated)
	err := row.Scan(card)
	if err != nil {
		return nil, errors.Wrap(err, "cardRepo.Create.QueryRowContext.Scan")
	}

	return card, nil
}

func (r *cardsRepo) GetByID(ctx context.Context, cardID int) (*models.CardModel, error) {
	card := &models.CardModel{}
	err := r.db.QueryRowContext(ctx, getCardById, cardID).Scan(card)
	if err != nil {
		return nil, errors.Wrap(err, "cardRepo.GetByID.QueryRowContext.Scan")
	}

	return card, nil
}

func (r *cardsRepo) GetAllCards(ctx context.Context, searchParams models.CardSearchParams, query *utilities.PaginationQuery) (*models.CardsList, error) {
	getAllCardsCountWithSearchParams := getTotalCountAllCards
	getAllCardsWithSearchParams := getAllCards

	var whereFilters []string
	if (searchParams != models.CardSearchParams{}) {
		getAllCardsCountWithSearchParams += " WHERE "
		getAllCardsCountWithSearchParams += " WHERE "
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
		getAllCardsCountWithSearchParams += strings.Join(whereFilters, " AND ")
		getAllCardsWithSearchParams += strings.Join(whereFilters, " AND ")
	}

	var totalRecords int
	err := r.db.QueryRowContext(ctx, getAllCardsCountWithSearchParams).Scan(totalRecords)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.GetAllCards.QueryRowContext")
	}
	if totalRecords == 0 {
		return &models.CardsList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]*models.CardModel, 0),
		}, nil
	}

	var cardsList []*models.CardModel
	rows, err := r.db.QueryContext(ctx, getAllCardsWithSearchParams)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.GetAllCards.QueryContext")
	}
	for rows.Next() {
		var card models.CardModel
		err := rows.Scan(&card)
		if err != nil {
			return nil, errors.Wrap(err, "cardsRepo.GetAllCards.QueryContext.Scan")
		}
		cardsList = append(cardsList, &card)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "cardsRepo.GetAllCards.rows.Err")
	}

	return &models.CardsList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         cardsList,
	}, nil
}

func (r *cardsRepo) Update(ctx context.Context, cardToUpdate *models.CardEntity) (*models.CardEntity, error) {
	syncDateUpdated := time.Now()

	card := &models.CardEntity{}
	err := r.db.QueryRowContext(ctx, updateCard, &cardToUpdate.MarketPrice, syncDateUpdated, &cardToUpdate.CardID).Scan(card)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.Update.QueryRowContext.Scan")
	}

	return card, nil
}

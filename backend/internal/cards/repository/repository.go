package repository

import (
	"backend/internal/cards"
	"backend/internal/models"
	"backend/pkg/utilities"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type cardsRepo struct {
	db *sqlx.DB
}

func NewCardsRepository(db *sqlx.DB) cards.Repository {
	return &cardsRepo{db: db}
}

func (r *cardsRepo) Create(ctx context.Context, newCard *models.CardModel) (*models.CardModel, error) {
	syncDateCreated := time.Now()
	syncDateUpdated := time.Now()

	card := &models.CardEntity{}
	row := r.db.QueryRowxContext(ctx, createCard, &newCard.CardCode, &newCard.CardName, &newCard.Set.SetID, &newCard.Supertype.SupertypeID, &newCard.Rarity.RarityID, &newCard.MarketPrice, &newCard.Pricetype.PricetypeID, &newCard.Image, syncDateCreated, syncDateUpdated)
	err := row.StructScan(card)
	if err != nil {
		return nil, errors.Wrap(err, "cardRepo.Create.QueryRowxContext.StructScan")
	}

	for _, t := range newCard.Types {
		newCardtype := &models.CardTypeModel{
			CardID: card.CardID,
			TypeID: t.TypeID,
		}
		_, err = r.CreateCardType(ctx, newCardtype)
		if err != nil {
			return nil, err
		}
	}

	for _, subtype := range newCard.Subtypes {
		newCardSubtype := &models.CardSubtypeModel{
			CardID:    card.CardID,
			SubtypeID: subtype.SubtypeID,
		}
		_, err = r.CreateCardSubtype(ctx, newCardSubtype)
		if err != nil {
			return nil, err
		}
	}

	return r.GetByID(ctx, card.CardID)
}

func (r *cardsRepo) GetByID(ctx context.Context, cardID int) (*models.CardModel, error) {
	card := &models.CardModel{}
	err := r.db.GetContext(ctx, card, getCardById, cardID)
	if err != nil {
		return nil, errors.Wrap(err, "cardRepo.GetByID.GetContext")
	}

	return card, nil
}

func (r *cardsRepo) Search(ctx context.Context, searchParams *models.CardSearchParams, query *utilities.PaginationQuery) (*models.CardsList, error) {
	var totalRecords int
	nstmt, err := r.db.PrepareNamedContext(ctx, utilities.FormatSqlQueryWithSearchParams(getTotalCountAllCards, *searchParams, false))
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.Search.PrepareNamedContext.getTotalCountAllCards")
	}
	err = nstmt.GetContext(ctx, &totalRecords, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.Search.GetContext")
	}
	if totalRecords == 0 {
		return &models.CardsList{
			TotalRecords: totalRecords,
			TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
			CurrentPage:  query.GetPage(),
			Size:         query.GetSize(),
			Data:         make([]models.CardModel, 0),
		}, nil
	}

	var cardsList []models.CardModel
	nstmt, err = r.db.PrepareNamedContext(ctx, fmt.Sprintf(utilities.FormatSqlQueryWithSearchParams(getAllCards, *searchParams, true), query.GetOffset(), query.GetLimit()))
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.Search.PrepareNamedContext.getAllCards")
	}
	err = nstmt.SelectContext(ctx, &cardsList, &searchParams)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.GetAllCards.SelectContext")
	}

	return &models.CardsList{
		TotalRecords: totalRecords,
		TotalPages:   utilities.GetTotalPages(totalRecords, query.GetSize()),
		CurrentPage:  query.GetPage(),
		Size:         query.GetSize(),
		Data:         cardsList,
	}, nil
}

func (r *cardsRepo) Update(ctx context.Context, cardID int, cardToUpdate *models.CardModel) (*models.CardModel, error) {
	syncDateUpdated := time.Now()

	card := &models.CardEntity{}
	row := r.db.QueryRowxContext(ctx, updateCard, &cardToUpdate.MarketPrice, syncDateUpdated, cardID)
	err := row.StructScan(card)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.Update.QueryRowxContext.StructScan")
	}

	return r.GetByID(ctx, card.CardID)
}

func (r *cardsRepo) AttachCardTypesAndSubtypes(ctx context.Context, card *models.CardModel) (*models.CardModel, error) {
	// TODO
	updatedCard := card
	cardtypes, err := r.GetCardTypes(ctx, card.CardID)
	if err != nil {
		return nil, err
	}
	updatedCard.Types = *cardtypes

	cardsubtypes, err := r.GetCardSubtypes(ctx, card.CardID)
	if err != nil {
		return nil, err
	}
	updatedCard.Subtypes = *cardsubtypes

	return updatedCard, nil
}

func (r *cardsRepo) CreateCardType(ctx context.Context, newCardtype *models.CardTypeModel) (*models.CardTypeModel, error) {
	cardtype := &models.CardTypeEntity{}
	row := r.db.QueryRowxContext(ctx, createCardType, &newCardtype.CardID, &newCardtype.TypeID)
	err := row.StructScan(cardtype)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.CreateCardType.QueryRowxContext.StructScan")
	}

	return &models.CardTypeModel{
		CardID: cardtype.CardID,
		TypeID: cardtype.TypeID,
	}, nil
}

func (r *cardsRepo) GetCardTypes(ctx context.Context, cardID int) (*[]models.TypeModel, error) {
	var cardtypesList *[]models.TypeModel
	err := r.db.SelectContext(ctx, cardtypesList, getCardtypesById, cardID)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.GetCardTypes.SelectContext")
	}

	return cardtypesList, nil
}

func (r *cardsRepo) CreateCardSubtype(ctx context.Context, newCardSubtype *models.CardSubtypeModel) (*models.CardSubtypeModel, error) {
	cardsubtype := &models.CardSubtypeEntity{}
	row := r.db.QueryRowxContext(ctx, createCardSubtype, &newCardSubtype.CardID, &newCardSubtype.SubtypeID)
	err := row.StructScan(cardsubtype)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.CreateCardSubtype.QueryRowxContext.StructScan")
	}

	return &models.CardSubtypeModel{
		CardID:    cardsubtype.CardID,
		SubtypeID: cardsubtype.SubtypeID,
	}, nil
}

func (r *cardsRepo) GetCardSubtypes(ctx context.Context, cardID int) (*[]models.SubtypeModel, error) {
	var cardsubtypesList *[]models.SubtypeModel
	err := r.db.SelectContext(ctx, cardsubtypesList, getCardsubtypesById, cardID)
	if err != nil {
		return nil, errors.Wrap(err, "cardsRepo.GetCardSubtypes.SelectContext")
	}

	return cardsubtypesList, nil
}

package main

import (
	"backend/config"
	cardsRepo "backend/internal/cards/repository"
	cardsService "backend/internal/cards/service"
	"backend/internal/models"
	pricetypesRepo "backend/internal/pricetypes/repository"
	pricetypesService "backend/internal/pricetypes/service"
	raritiesRepo "backend/internal/rarities/repository"
	raritiesService "backend/internal/rarities/service"
	seriesRepo "backend/internal/series/repository"
	seriesService "backend/internal/series/service"
	setsRepo "backend/internal/sets/repository"
	setsService "backend/internal/sets/service"
	subtypesRepo "backend/internal/subtypes/repository"
	subtypesService "backend/internal/subtypes/service"
	supertypesRepo "backend/internal/supertypes/repository"
	supertypesService "backend/internal/supertypes/service"
	typesRepo "backend/internal/types/repository"
	typesService "backend/internal/types/service"
	"backend/pkg/db"
	"backend/pkg/utilities"
	"context"
	"log"
	"os"
	"strings"
	"time"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"

	_ "github.com/lib/pq"
)

var fullSyncFlag, halfSyncFlag = false, false

func main() {
	configFile, err := config.OpenConfig(os.Getenv("config"))
	if err != nil {
		log.Fatalf("Could not open config: %v", err)
	}

	config, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Could not parse config: %v", err)
	}

	db, err := db.NewPsqlDB(config)
	if err != nil {
		log.Fatalf("Could not connect to Datbase: %v", err)
	}

	// Init repositories
	cardsRepo := cardsRepo.NewCardsRepository(db)
	raritiesRepo := raritiesRepo.NewRaritiesRepository(db)
	seriesRepo := seriesRepo.NewSeriesRepository(db)
	setsRepo := setsRepo.NewSetsRepository(db)
	subtypesRepo := subtypesRepo.NewSubtypesRepository(db)
	supertypesRepo := supertypesRepo.NewSupertypesRepository(db)
	typesRepo := typesRepo.NewTypesRepository(db)
	pricetypesRepo := pricetypesRepo.NewPricetypesRepository(db)

	// Init services
	cardsService := cardsService.NewCardsService(cardsRepo)
	raritiesService := raritiesService.NewRaritiesService(raritiesRepo)
	seriesService := seriesService.NewSeriesService(seriesRepo)
	setsService := setsService.NewSeriesService(setsRepo)
	subtypesService := subtypesService.NewSubtypesService(subtypesRepo)
	supertypesService := supertypesService.NewSupertypesService(supertypesRepo)
	typesService := typesService.NewTypesService(typesRepo)
	pricetypesService := pricetypesService.NewPricetypesService(pricetypesRepo)

	args := os.Args[1:]
	if len(args) == 1 && args[0] == "full" {
		fullSyncFlag = true
	} else if len(args) == 1 && args[0] == "half" {
		halfSyncFlag = true
	}

	tcgClient := tcg.NewClient(config.TCGApiKey)

	ctx := context.Background()

	log.Printf("Sync initiated at: %s", time.Now().String())

	if halfSyncFlag {
		sets, err := tcgClient.GetSets()
		if err != nil {
			log.Fatalf("Error getting sets from API: %v", err)
		}
		for _, curSet := range sets {
			// Check if the set is in the database
			var setLookup *models.SetsList
			setSearchParams := models.SetSearchParams{
				SetCode: curSet.ID,
			}
			setLookup, err := setsService.Search(ctx, &setSearchParams, utilities.NewPaginationQuery(1, 1))
			if err != nil {
				log.Fatalf("Error searching sets: %v", err)
			}
			if setLookup.TotalRecords == 0 {
				// If set does not exist, check if the series exists
				var series *models.SeriesModel
				var seriesLookup *models.SeriesList
				seriesSearchParams := models.SeriesSearchParams{
					SeriesName: curSet.Series,
				}
				seriesLookup, err := seriesService.Search(ctx, &seriesSearchParams, utilities.NewPaginationQuery(1, 1))
				if err != nil {
					log.Fatalf("Error searching series: %v", err)
				}
				if seriesLookup.TotalRecords == 0 {
					// If series does not exist, create the series
					newSeries := models.SeriesModel{SeriesName: curSet.Series}
					series, err = seriesService.Create(ctx, &newSeries)
					if err != nil {
						log.Fatalf("Error creating Series: %v", err)
					}
					log.Printf("Created new Series: %s", newSeries.SeriesName)
				} else {
					series = &seriesLookup.Data[0]
				}
				// Create the set
				formattedReleaseDate := strings.ReplaceAll(curSet.ReleaseDate, "/", "-")
				setReleaseDate, err := time.Parse(time.DateOnly, formattedReleaseDate)
				if err != nil {
					log.Fatalf("Error parsing time: %v", err)
				}
				newSet := models.SetModel{
					SetName:           curSet.Name,
					SetCode:           curSet.ID,
					Series:            *series,
					PtcgoCode:         curSet.PtcgoCode,
					CardTotal:         curSet.PrintedTotal,
					ExtendedCardTotal: curSet.Total,
					SetReleaseDate:    setReleaseDate,
					SymbolImage:       curSet.Images.Symbol,
					LogoImage:         curSet.Images.Logo,
				}
				createdSet, err := setsService.Create(ctx, &newSet)
				if err != nil {
					log.Fatalf("Error creating set: %v", err)
				}
				log.Printf("Created new Set: %s - %s", createdSet.SetCode, createdSet.SetName)
			}
		}
		log.Printf("Half Sync Complete at: %s", time.Now().String())
	}
	if fullSyncFlag {
		// Sync Types
		types, err := tcgClient.GetTypes()
		if err != nil {
			log.Fatalf("Error getting types from API: %v", err)
		}
		for _, curType := range types {
			var typeLookup *models.TypesList
			typeSearchParams := models.TypeSearchParams{
				TypeName: curType,
			}
			typeLookup, err := typesService.Search(ctx, &typeSearchParams, utilities.NewPaginationQuery(1, 1))
			if err != nil {
				log.Fatalf("Error searching types %v", err)
			}
			if typeLookup.TotalRecords == 0 {
				// If type does not exist, create the type
				newType := models.TypeModel{
					TypeName: curType,
				}
				createdType, err := typesService.Create(ctx, &newType)
				if err != nil {
					log.Fatalf("Error creating Type: %v", err)
				}
				log.Printf("Created new Type: %s", createdType.TypeName)
			}
		}
		// Sync Subtypes
		subtypes, err := tcgClient.GetSubTypes()
		if err != nil {
			log.Fatalf("Error getting subtypes from API: %v", err)
		}
		for _, curSubtype := range subtypes {
			var subtypeLookup *models.SubtypesList
			subtypeSearchParams := models.SubtypeSearchParams{
				SubtypeName: curSubtype,
			}
			subtypeLookup, err := subtypesService.Search(ctx, &subtypeSearchParams, utilities.NewPaginationQuery(1, 1))
			if err != nil {
				log.Fatalf("Error searching subtypes %v", err)
			}
			if subtypeLookup.TotalRecords == 0 {
				// If type does not exist, create the type
				newSubtype := models.SubtypeModel{
					SubtypeName: curSubtype,
				}
				createdSubtype, err := subtypesService.Create(ctx, &newSubtype)
				if err != nil {
					log.Fatalf("Error creating Subtype: %v", err)
				}
				log.Printf("Created new Subtype: %s", createdSubtype.SubtypeName)
			}
		}
		// Sync Supertypes
		supertypes, err := tcgClient.GetSuperTypes()
		if err != nil {
			log.Fatalf("Error getting supertypes from API: %v", err)
		}
		for _, curSupertype := range supertypes {
			var supertypeLookup *models.SupertypesList
			supertypeSearchParams := models.SupertypeSearchParams{
				SupertypeName: curSupertype,
			}
			supertypeLookup, err := supertypesService.Search(ctx, &supertypeSearchParams, utilities.NewPaginationQuery(1, 1))
			if err != nil {
				log.Fatalf("Error searching supertype %v", err)
			}
			if supertypeLookup.TotalRecords == 0 {
				// If type does not exist, create the type
				newsupertype := models.SupertypeModel{
					SupertypeName: curSupertype,
				}
				createdsupertype, err := supertypesService.Create(ctx, &newsupertype)
				if err != nil {
					log.Fatalf("Error creating Supertype: %v", err)
				}
				log.Printf("Created new Supertype: %s", createdsupertype.SupertypeName)
			}
		}
		// Sync Rarities
		rarities, err := tcgClient.GetRarities()
		if err != nil {
			log.Fatalf("Error getting rarities from API: %v", err)
		}
		for _, curRarity := range rarities {
			var rarityLookup *models.RaritiesList
			raritiesSearchParams := models.RaritySearchParams{
				RarityName: curRarity,
			}
			rarityLookup, err := raritiesService.Search(ctx, &raritiesSearchParams, utilities.NewPaginationQuery(1, 1))
			if err != nil {
				log.Fatalf("Error searching rarity: %v", err)
			}
			if rarityLookup.TotalRecords == 0 {
				// If type does not exist, create the type
				newRarity := models.RarityModel{
					RarityName: curRarity,
				}
				createdRarity, err := raritiesService.Create(ctx, &newRarity)
				if err != nil {
					log.Fatalf("Error creating Rarity: %v", err)
				}
				log.Printf("Created new Rarity: %s", createdRarity.RarityName)
			}
		}
		// Pricetypes will remain constant
		log.Printf("Full Sync Complete at: %s", time.Now().String())
	}
	page := 0
	cards, err := tcgClient.GetCards()
	if err != nil {
		log.Fatalf("Error getting cards from API on page: %d", page)
	}
	for len(cards) > 0 {
		for _, curCard := range cards {
			var normalPrice, holoFoilPrice, reverseHolofoilPrice float32
			var curPricetype string
			if curCard.TCGPlayer.Prices.Normal != nil {
				normalPrice = float32(curCard.TCGPlayer.Prices.Normal.Market)
			}
			if curCard.TCGPlayer.Prices.Holofoil != nil {
				holoFoilPrice = float32(curCard.TCGPlayer.Prices.Holofoil.Market)
			}
			if curCard.TCGPlayer.Prices.ReverseHolofoil != nil {
				reverseHolofoilPrice = float32(curCard.TCGPlayer.Prices.ReverseHolofoil.Market)
			}

			// Get set model from db
			var setLookup *models.SetsList
			var set *models.SetModel
			setSearchParams := models.SetSearchParams{
				SetName: curCard.Set.Name,
			}
			setLookup, err = setsService.Search(ctx, &setSearchParams, utilities.NewPaginationQuery(1, 1))
			if err != nil {
				log.Fatalf("Error searching sets: %v", err)
			}
			if setLookup.TotalRecords == 0 {
				log.Printf("Card set %s does not exist - rerun half sync", curCard.Set.Name)
				continue
			} else {
				set = &setLookup.Data[0]
			}
			// Get supertype model from db
			var supertypeLookup *models.SupertypesList
			var supertype *models.SupertypeModel
			supertypeSearchParams := models.SupertypeSearchParams{
				SupertypeName: curCard.Supertype,
			}
			supertypeLookup, err = supertypesService.Search(ctx, &supertypeSearchParams, utilities.NewPaginationQuery(1, 1))
			if err != nil {
				log.Fatalf("Error searching supertypes: %v", err)
			}
			if supertypeLookup.TotalRecords == 0 {
				log.Printf("Card supertype %s does not exist - rerun full sync", curCard.Supertype)
				continue
			} else {
				supertype = &supertypeLookup.Data[0]
			}
			// Get rarity model from db
			var rarityLookup *models.RaritiesList
			var rarity *models.RarityModel
			raritySearchParams := models.RaritySearchParams{
				RarityName: curCard.Rarity,
			}
			rarityLookup, err = raritiesService.Search(ctx, &raritySearchParams, utilities.NewPaginationQuery(1, 1))
			if err != nil {
				log.Fatalf("Error searching rarities: %v", err)
			}
			if rarityLookup.TotalRecords == 0 {
				log.Printf("Card rarity %s does not exist - rerun full sync", curCard.Rarity)
				continue
			} else {
				rarity = &rarityLookup.Data[0]
			}
			// Get type models from db
			var typesList []models.TypeModel
			var typeLookup *models.TypesList
			for _, curType := range curCard.Types {
				typeSearchParams := models.TypeSearchParams{
					TypeName: curType,
				}
				typeLookup, err = typesService.Search(ctx, &typeSearchParams, utilities.NewPaginationQuery(1, 1))
				if err != nil {
					log.Fatalf("Error searching types: %v", err)
				}
				typesList = append(typesList, typeLookup.Data...)
			}
			// Get subtype models from db
			var subtypesList []models.SubtypeModel
			var subtypeLookup *models.SubtypesList
			for _, curSubtype := range curCard.Subtypes {
				subtypeSearchParams := models.SubtypeSearchParams{
					SubtypeName: curSubtype,
				}
				subtypeLookup, err = subtypesService.Search(ctx, &subtypeSearchParams, utilities.NewPaginationQuery(1, 1))
				if err != nil {
					log.Fatalf("Error searching subtypes: %v", err)
				}
				subtypesList = append(subtypesList, subtypeLookup.Data...)
			}

			// Normal Card
			if normalPrice != 0.0 {
				curPricetype = "Normal"
				var cardLookup *models.CardsList
				cardSearchParams := models.CardSearchParams{
					CardCode:      curCard.ID,
					PricetypeName: curPricetype,
				}
				cardLookup, err := cardsService.Search(ctx, &cardSearchParams, utilities.NewPaginationQuery(1, 1))
				if err != nil {
					log.Fatalf("Error searching card: %v", err)
				}
				if cardLookup.TotalRecords == 0 {
					// If card does not exist, create the card
					// Get pricetype model from db
					var pricetypeLookup *models.PricetypesList
					var pricetype *models.PricetypeModel
					pricetypeSearchParams := models.PricetypeSearchParams{
						PricetypeName: curPricetype,
					}
					pricetypeLookup, err = pricetypesService.Search(ctx, &pricetypeSearchParams, utilities.NewPaginationQuery(1, 1))
					if err != nil {
						log.Fatalf("Error searching pricetypes: %v", err)
					}
					if pricetypeLookup.TotalRecords == 0 {
						log.Printf("Card pricetype %s does not exist - rerun full sync", curPricetype)
						continue
					} else {
						pricetype = &pricetypeLookup.Data[0]
					}
					newCard := models.CardModel{
						CardName:    curCard.Name,
						CardCode:    curCard.ID,
						Set:         *set,
						Supertype:   *supertype,
						Types:       typesList,
						Subtypes:    subtypesList,
						Rarity:      *rarity,
						MarketPrice: normalPrice,
						Pricetype:   *pricetype,
						Image:       curCard.Images.Large,
					}
					createdCard, err := cardsService.Create(ctx, &newCard)
					if err != nil {
						log.Fatalf("Error creating Card: %v", err)
					}
					log.Printf("Created new Card: %s - %s - %s", createdCard.CardCode, curPricetype, createdCard.CardName)
				}
			}

			// Holofoil Card
			if holoFoilPrice != 0.0 {
				curPricetype = "Holofoil"
				var cardLookup *models.CardsList
				cardSearchParams := models.CardSearchParams{
					CardCode:      curCard.ID,
					PricetypeName: curPricetype,
				}
				cardLookup, err := cardsService.Search(ctx, &cardSearchParams, utilities.NewPaginationQuery(1, 1))
				if err != nil {
					log.Fatalf("Error searching card: %v", err)
				}
				if cardLookup.TotalRecords == 0 {
					// If card does not exist, create the card
					// Get pricetype model from db
					var pricetypeLookup *models.PricetypesList
					var pricetype *models.PricetypeModel
					pricetypeSearchParams := models.PricetypeSearchParams{
						PricetypeName: curPricetype,
					}
					pricetypeLookup, err = pricetypesService.Search(ctx, &pricetypeSearchParams, utilities.NewPaginationQuery(1, 1))
					if err != nil {
						log.Fatalf("Error searching pricetypes: %v", err)
					}
					if pricetypeLookup.TotalRecords == 0 {
						log.Printf("Card pricetype %s does not exist - rerun full sync", curPricetype)
						continue
					} else {
						pricetype = &pricetypeLookup.Data[0]
					}
					newCard := models.CardModel{
						CardName:    curCard.Name,
						CardCode:    curCard.ID,
						Set:         *set,
						Supertype:   *supertype,
						Types:       typesList,
						Subtypes:    subtypesList,
						Rarity:      *rarity,
						MarketPrice: normalPrice,
						Pricetype:   *pricetype,
						Image:       curCard.Images.Large,
					}
					createdCard, err := cardsService.Create(ctx, &newCard)
					if err != nil {
						log.Fatalf("Error creating Card: %v", err)
					}
					log.Printf("Created new Card: %s - %s - %s", createdCard.CardCode, curPricetype, createdCard.CardName)
				}
			}

			// ReverseHolofoil Card
			if reverseHolofoilPrice != 0.0 {
				curPricetype = "ReverseHolofoil"
				var cardLookup *models.CardsList
				cardSearchParams := models.CardSearchParams{
					CardCode:      curCard.ID,
					PricetypeName: curPricetype,
				}
				cardLookup, err := cardsService.Search(ctx, &cardSearchParams, utilities.NewPaginationQuery(1, 1))
				if err != nil {
					log.Fatalf("Error searching card: %v", err)
				}
				if cardLookup.TotalRecords == 0 {
					// If card does not exist, create the card
					// Get pricetype model from db
					var pricetypeLookup *models.PricetypesList
					var pricetype *models.PricetypeModel
					pricetypeSearchParams := models.PricetypeSearchParams{
						PricetypeName: curPricetype,
					}
					pricetypeLookup, err = pricetypesService.Search(ctx, &pricetypeSearchParams, utilities.NewPaginationQuery(1, 1))
					if err != nil {
						log.Fatalf("Error searching pricetypes: %v", err)
					}
					if pricetypeLookup.TotalRecords == 0 {
						log.Printf("Card pricetype %s does not exist - rerun full sync", curPricetype)
						continue
					} else {
						pricetype = &pricetypeLookup.Data[0]
					}
					newCard := models.CardModel{
						CardName:    curCard.Name,
						CardCode:    curCard.ID,
						Set:         *set,
						Supertype:   *supertype,
						Types:       typesList,
						Subtypes:    subtypesList,
						Rarity:      *rarity,
						MarketPrice: normalPrice,
						Pricetype:   *pricetype,
						Image:       curCard.Images.Large,
					}
					createdCard, err := cardsService.Create(ctx, &newCard)
					if err != nil {
						log.Fatalf("Error creating Card: %v", err)
					}
					log.Printf("Created new Card: %s - %s - %s", createdCard.CardCode, curPricetype, createdCard.CardName)
				}
			}
		}

		page += 1
		cards, err = tcgClient.GetCards(request.Page(page))
		if err != nil {
			log.Printf("Error getting cards on page: %d", page)
			log.Fatal(err)
		}
	}

	log.Printf("Sync Complete at: %s", time.Now().String())
}

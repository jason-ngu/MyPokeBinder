package main

import (
	common "backend/common"
	internal "backend/internal"
	models "backend/internal/models"
	cardsService "backend/internal/services/cards"
	raritiesService "backend/internal/services/rarities"
	seriesService "backend/internal/services/series"
	setsService "backend/internal/services/sets"
	subtypesService "backend/internal/services/subtypes"
	superTypesService "backend/internal/services/supertypes"
	typesService "backend/internal/services/types"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
	_ "github.com/lib/pq"
)

var fullSyncFlag, halfSyncFlag = false, false
var timeFormat = time.RFC3339

func main() {
	config := common.SetupConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DatabaseName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	env := &internal.Env{DB: db}

	tcgClient := tcg.NewClient(config.ApiKey)

	args := os.Args[1:]

	if len(args) == 1 && args[0] == "full" {
		fullSyncFlag = true
	} else if len(args) == 1 && args[0] == "half" {
		halfSyncFlag = true
	}

	log.Printf("Sync initiated at: %s", time.Now().String())

	if halfSyncFlag {
		sets, err := tcgClient.GetSets()
		if err != nil {
			log.Printf("Error getting sets")
			log.Fatal(err)
		}
		for _, curSet := range sets {
			// Check if series exists
			var seriesLookup models.SeriesModel
			seriesLookup, err := seriesService.GetSeriesByName(env, curSet.Series)
			if (err != nil) && (seriesLookup == models.SeriesModel{}) {
				// If it does not, create it in the database
				newSeries := models.SeriesModel{SeriesName: curSet.Series}
				seriesLookup, err = seriesService.CreateSeries(env, newSeries)
				if err != nil {
					log.Printf("Error creating series: %s", curSet.Series)
					log.Fatal(err)
				}
				log.Printf("Created new Series: %s", newSeries.SeriesName)
			}

			// Create a new set
			setReleaseDate, err := time.Parse(timeFormat, curSet.ReleaseDate)
			if err != nil {
				log.Fatal(err)
			}

			newSet := models.SetModel{
				SetName:           curSet.Name,
				SetCode:           curSet.ID,
				SeriesName:        seriesLookup.SeriesName,
				PtcgoCode:         curSet.PtcgoCode,
				CardTotal:         curSet.PrintedTotal,
				ExtendedCardTotal: curSet.Total,
				SetReleaseDate:    setReleaseDate,
				SymbolImage:       curSet.Images.Symbol,
				LogoImage:         curSet.Images.Logo,
			}
			createdSet, err := setsService.CreateSet(env, newSet)
			if err != nil {
				log.Printf("Error creating set: %s - %s - %s", curSet.ID, curSet.Name, seriesLookup.SeriesName)
				log.Fatal(err)
			}
			log.Printf("Created new Set: %s", createdSet.SetName)
		}
		log.Printf("Half Sync Complete at: %s", time.Now().String())
	}
	if fullSyncFlag {
		// Sync Types to DB
		types, err := tcgClient.GetTypes()
		if err != nil {
			log.Printf("Error getting types")
			log.Fatal(err)
		}
		for _, curType := range types {
			newType := models.TypeModel{TypeName: curType}
			createdType, err := typesService.CreateType(env, newType)
			if err != nil {
				log.Printf("Error creating type: %s", curType)
				log.Fatal(err)
			}
			log.Printf("Created new Type: %s", createdType.TypeName)
		}
		// Sync Subtypes to DB
		subtypes, err := tcgClient.GetSubTypes()
		if err != nil {
			log.Printf("Error getting subtypes")
			log.Fatal(err)
		}
		for _, curSubtype := range subtypes {
			newSubtype := models.SubtypeModel{SubtypeName: curSubtype}
			createdSubtype, err := subtypesService.CreateSubtype(env, newSubtype)
			if err != nil {
				log.Printf("Error creating supertype: %s", curSubtype)
				log.Fatal(err)
			}
			log.Printf("Created new Subtype: %s", createdSubtype.SubtypeName)
		}
		// Sync Supertypes to DB
		supertypes, err := tcgClient.GetSuperTypes()
		if err != nil {
			log.Printf("Error getting supertypes")
			log.Fatal(err)
		}
		for _, curSupertype := range supertypes {
			newSupertype := models.SupertypeModel{SupertypeName: curSupertype}
			createdSupertype, err := superTypesService.CreateSupertype(env, newSupertype)
			if err != nil {
				log.Printf("Error creating supertype: %s", curSupertype)
				log.Fatal(err)
			}
			log.Printf("Created new Supertype: %s", createdSupertype.SupertypeName)
		}
		// Sync Rarities to DB
		rarities, err := tcgClient.GetRarities()
		if err != nil {
			log.Printf("Error getting rarities")
			log.Fatal(err)
		}
		for _, curRarity := range rarities {
			newRarity := models.RarityModel{RarityName: curRarity}
			createdRarity, err := raritiesService.CreateRarity(env, newRarity)
			if err != nil {
				log.Printf("Error creating rarity: %s", curRarity)
				log.Fatal(err)
			}
			log.Printf("Created new Rarity: %s", createdRarity.RarityName)
		}
		log.Printf("Full Sync Complete at: %s", time.Now().String())
		// PriceTypes will remain constant
	}

	page := 0
	cards, err := tcgClient.GetCards()
	if err != nil {
		log.Printf("Error getting cards on page: %d", page)
		log.Fatal(err)
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

			// Normal Card
			if normalPrice != 0.0 {
				curPricetype = "Normal"
				cardLookup, err := cardsService.GetCard(env, curCard.ID, curPricetype)
				if err != nil {
					log.Printf("Error finding card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
					log.Fatal(err)
				}

				if (cardLookup != models.CardModel{}) {
					// If the card exists, update it
					cardToUpdate := models.CardModel{
						CardName:      curCard.Name,
						CardCode:      curCard.ID,
						MarketPrice:   normalPrice,
						PricetypeName: curPricetype,
					}
					_, err := cardsService.UpdateCard(env, cardToUpdate)
					if err != nil {
						log.Printf("Error updating card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
						log.Fatal(err)
					}
					log.Printf("Successfully updated card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
				} else {
					// Otherwise create a new card
					newCard := models.CardModel{
						CardName:      curCard.Name,
						CardCode:      curCard.ID,
						SetName:       curCard.Set.Name,
						SupertypeName: curCard.Supertype,
						RarityName:    curCard.Rarity,
						MarketPrice:   normalPrice,
						PricetypeName: curPricetype,
						Image:         curCard.Images.Large,
					}
					_, err := cardsService.CreateCard(env, newCard)
					if err != nil {
						log.Printf("Error creating card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
						log.Fatal(err)
					}
					log.Printf("Successfully created card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
				}
			}

			// Holofoil Card
			if holoFoilPrice != 0.0 {
				curPricetype = "Holofoil"
				cardLookup, err := cardsService.GetCard(env, curCard.ID, curPricetype)
				if err != nil {
					log.Printf("Error finding card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
					log.Fatal(err)
				}

				if (cardLookup != models.CardModel{}) {
					// If the card exists, update it
					cardToUpdate := models.CardModel{
						CardName:      curCard.Name,
						CardCode:      curCard.ID,
						MarketPrice:   holoFoilPrice,
						PricetypeName: curPricetype,
					}
					_, err := cardsService.UpdateCard(env, cardToUpdate)
					if err != nil {
						log.Printf("Error updating card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
						log.Fatal(err)
					}
					log.Printf("Successfully updated card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
				} else {
					// Otherwise create a new card
					newCard := models.CardModel{
						CardName:      curCard.Name,
						CardCode:      curCard.ID,
						SetName:       curCard.Set.Name,
						SupertypeName: curCard.Supertype,
						RarityName:    curCard.Rarity,
						MarketPrice:   holoFoilPrice,
						PricetypeName: curPricetype,
						Image:         curCard.Images.Large,
					}
					_, err := cardsService.CreateCard(env, newCard)
					if err != nil {
						log.Printf("Error creating card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
						log.Fatal(err)
					}
					log.Printf("Successfully created card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
				}
			}

			// ReverseHolofoil Card
			if reverseHolofoilPrice != 0.0 {
				curPricetype = "ReverseHolofoil"
				cardLookup, err := cardsService.GetCard(env, curCard.ID, curPricetype)
				if err != nil {
					log.Printf("Error finding card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
					log.Fatal(err)
				}

				if (cardLookup != models.CardModel{}) {
					// If the card exists, update it
					cardToUpdate := models.CardModel{
						CardName:      curCard.Name,
						CardCode:      curCard.ID,
						MarketPrice:   reverseHolofoilPrice,
						PricetypeName: curPricetype,
					}
					_, err := cardsService.UpdateCard(env, cardToUpdate)
					if err != nil {
						log.Printf("Error updating card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
						log.Fatal(err)
					}
					log.Printf("Successfully updated card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
				} else {
					// Otherwise create a new card
					newCard := models.CardModel{
						CardName:      curCard.Name,
						CardCode:      curCard.ID,
						SetName:       curCard.Set.Name,
						SupertypeName: curCard.Supertype,
						RarityName:    curCard.Rarity,
						MarketPrice:   reverseHolofoilPrice,
						PricetypeName: curPricetype,
						Image:         curCard.Images.Large,
					}
					_, err := cardsService.CreateCard(env, newCard)
					if err != nil {
						log.Printf("Error creating card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
						log.Fatal(err)
					}
					log.Printf("Successfully created card %s - %s - %s", curCard.ID, curCard.Name, curPricetype)
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

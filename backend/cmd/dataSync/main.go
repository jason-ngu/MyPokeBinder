package main

import (
	common "backend/common"
	internal "backend/internal"
	models "backend/internal/models"
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

	if halfSyncFlag {
		sets, err := tcgClient.GetSets()
		if err != nil {
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
					log.Fatal(err)
				}
				log.Printf("Created new Series: %s", newSeries.SeriesName)
			}

			// Create a new set
			setReleaseDate, err := time.Parse(timeFormat, curSet.ReleaseDate)

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
				log.Fatal(err)
			}
			log.Printf("Created new Set: %s", createdSet.SetName)
		}
		log.Println("Half Sync Complete")
	}
	if fullSyncFlag {
		// Sync Types to DB
		types, err := tcgClient.GetTypes()
		if err != nil {
			log.Fatal(err)
		}
		for _, curType := range types {
			newType := models.TypeModel{TypeName: curType}
			createdType, err := typesService.CreateType(env, newType)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Created new Type: %s", createdType.TypeName)
		}
		// Sync Subtypes to DB
		subtypes, err := tcgClient.GetSubTypes()
		if err != nil {
			log.Fatal(err)
		}
		for _, curSubtype := range subtypes {
			newSubtype := models.SubtypeModel{SubtypeName: curSubtype}
			createdSubtype, err := subtypesService.CreateSubtype(env, newSubtype)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Created new Subtype: %s", createdSubtype.SubtypeName)
		}
		// Sync Supertypes to DB
		supertypes, err := tcgClient.GetSuperTypes()
		if err != nil {
			log.Fatal(err)
		}
		for _, curSupertype := range supertypes {
			newSupertype := models.SupertypeModel{SupertypeName: curSupertype}
			createdSupertype, err := superTypesService.CreateSupertype(env, newSupertype)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Created new Supertype: %s", createdSupertype.SupertypeName)
		}
		// Sync Rarities to DB
		rarities, err := tcgClient.GetRarities()
		if err != nil {
			log.Fatal(err)
		}
		for _, curRarity := range rarities {
			newRarity := models.RarityModel{RarityName: curRarity}
			createdRarity, err := raritiesService.CreateRarity(env, newRarity)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Created new Rarity: %s", createdRarity.RarityName)
		}
		log.Println("Full Sync Complete")
		// PriceTypes will remain constant
	}
	// Sync Cards and Price data to DB
	log.Println("Sync Complete")
}

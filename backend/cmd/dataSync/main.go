package main

import (
	"backend/config"
	"backend/internal/models"
	seriesRepo "backend/internal/series/repository"
	seriesService "backend/internal/series/service"
	setsService "backend/internal/sets/service"
	"context"
	"strings"
	"time"

	// cardsService "backend/internal/services/cards"
	// raritiesService "backend/internal/services/rarities"
	// seriesService "backend/internal/services/series"
	// setsService "backend/internal/services/sets"
	// subtypesService "backend/internal/services/subtypes"
	// superTypesService "backend/internal/services/supertypes"
	// typesService "backend/internal/services/types"
	setsRepo "backend/internal/sets/repository"
	"backend/pkg/db"
	"backend/pkg/utilities"
	"log"
	"os"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"

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
	// cardsRepo := cardsRepo.NewCardsRepository(db)
	// raritiesRepo := raritiesRepo.NewRaritiesRepository(db)
	seriesRepo := seriesRepo.NewSeriesRepository(db)
	setsRepo := setsRepo.NewSetsRepository(db)
	// subtypesRepo := subtypesRepo.NewSubtypesRepository(db)
	// supertypesRepo := supertypesRepo.NewSupertypesRepository(db)
	// typesRepo := typesRepo.NewTypesRepository(db)

	// Init services
	// cardsService := cardsService.NewCardsService(cardsRepo)
	// raritiesService := raritiesService.NewRaritiesService(raritiesRepo)
	seriesService := seriesService.NewSeriesService(seriesRepo)
	setsService := setsService.NewSeriesService(setsRepo)
	// subtypesService := subtypesService.NewSubtypesService(subtypesRepo)
	// superTypesService := superTypesService.NewSupertypesService(supertypesRepo)
	// typesService := typesService.NewTypesService(typesRepo)

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
			setLookup, err := setsService.SearchSets(ctx, &setSearchParams, utilities.NewPaginationQuery(10, 1))
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
				seriesLookup, err := seriesService.SearchSeries(ctx, &seriesSearchParams, utilities.NewPaginationQuery(10, 1))
				if err != nil {
					log.Fatalf("Error searching series: %v", err)
				}
				if seriesLookup.TotalRecords == 0 {
					// If series does not exist, create the series
					newSeries := models.SeriesModel{SeriesName: curSet.Series}
					series, err = seriesService.Create(ctx, &newSeries)
					if err != nil {
						log.Fatalf("Error creating series: %v", err)
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
}

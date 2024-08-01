package main

import (
	common "backend/common"
	internal "backend/internal"
	models "backend/internal/models"
	raritiesService "backend/internal/services/rarities"
	seriesService "backend/internal/services/series"
	subtypesService "backend/internal/services/subtypes"
	superTypesService "backend/internal/services/supertypes"
	typesService "backend/internal/services/types"
	"database/sql"
	"fmt"
	"log"
	"os"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	_ "github.com/lib/pq"
)

var fullSyncFlag, halfSyncFlag = false, false

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
			newSeries := models.SeriesModel{SeriesName: curSet.Series}
			seriesService.CreateSeries(env, newSeries)
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
			newTypeId, err := typesService.CreateType(env, newType)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Created new Type: %s with id: %d", curType, newTypeId)
		}
		// Sync Subtypes to DB
		subtypes, err := tcgClient.GetSubTypes()
		if err != nil {
			log.Fatal(err)
		}
		for _, curSubtype := range subtypes {
			newSubtype := models.SubtypeModel{SubtypeName: curSubtype}
			newSubtypeId, err := subtypesService.CreateSubtype(env, newSubtype)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Created new Subtype: %s with id: %d", curSubtype, newSubtypeId)
		}
		// Sync Supertypes to DB
		supertypes, err := tcgClient.GetSuperTypes()
		if err != nil {
			log.Fatal(err)
		}
		for _, curSupertype := range supertypes {
			newSupertype := models.SupertypeModel{SupertypeName: curSupertype}
			newSupertypeId, err := superTypesService.CreateSupertype(env, newSupertype)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Created new Supertype: %s with id: %d", curSupertype, newSupertypeId)
		}
		// Sync Rarities to DB
		rarities, err := tcgClient.GetRarities()
		if err != nil {
			log.Fatal(err)
		}
		for _, curRarity := range rarities {
			newRarity := models.RarityModel{RarityName: curRarity}
			newRarityId, err := raritiesService.CreateRarity(env, newRarity)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Created new Rarity: %s with id: %d", curRarity, newRarityId)
		}
		log.Println("Full Sync Complete")
		// PriceTypes will remain constant
		// Series will remain constant
	}
	// Sync Cards and Price data to DB
	log.Println("Sync Complete")
}

// type Series struct {
// 	series_name string
// }

// type Sets struct {
// 	api_id              string
// 	set_name            string
// 	series_id           int
// 	ptcgo_code          string
// 	card_total          string
// 	extended_card_total string
// 	set_release_date    time.Time
// 	symbol_image        string
// 	logo_image          string
// 	sync_date_created   time.Time
// 	sync_date_updated   time.Time
// }

// var fullSyncFlag, halfSyncFlag = false, false
// var db *sql.DB
// var DB *sqlx.DB

// func doesDataExist(table string, column string, dataLookupValue string) bool {
// 	cleanedData := cleanData(dataLookupValue)
// 	sqlStatement := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = '%s'", table, column, cleanedData)
// 	var foundRows bool
// 	err := DB.QueryRow(sqlStatement).Scan(&foundRows)
// 	log.Println(foundRows)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return foundRows
// }

// func cleanData(data string) string {
// 	return strings.ReplaceAll(data, "'", "''")
// }

// func getData(tableObj interface{}, lookupFields []string) interface{} {
// 	v := reflect.ValueOf(tableObj)
// 	tableName := v.Type().Name()

// 	sqlStatement := fmt.Sprintf("SELECT * FROM %s", tableName)
// 	log.Println(sqlStatement)
// 	var returnedEntries []interface{}
// 	err := DB.QueryRow(sqlStatement).Scan(&returnedEntries)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("New Record is: ", returnedEntries)
// 	return returnedEntries
// }

// func searchData(tableObj interface{}) []interface{} {
// 	v := reflect.ValueOf(tableObj)
// 	tableName := v.Type().Name()

// 	sqlStatement := fmt.Sprintf("SELECT * FROM %s", tableName)
// 	log.Println(sqlStatement)
// 	var returnedEntries []interface{}
// 	err := DB.QueryRow(sqlStatement).Scan(&returnedEntries)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("New Record is: ", returnedEntries)
// 	return returnedEntries
// }

// func insertNewData1(tableObj interface{}, idToReturn string) interface{} {
// 	v := reflect.ValueOf(tableObj)

// 	tableName := v.Type().Name()

// 	var fieldNames []string
// 	var fieldValues []string

// 	for i := 0; i < v.NumField(); i++ {
// 		// Get the field value
// 		fieldValue := v.Field(i)
// 		// Get the field name
// 		fieldName := v.Type().Field(i).Name
// 		// Print the field name and value
// 		fieldNames = append(fieldNames, fieldName)
// 		value := fmt.Sprintf("'%s'", cleanData(fieldValue.String()))
// 		fieldValues = append(fieldValues, value)
// 	}

// 	formattedColumns := strings.Join(fieldNames, ", ")
// 	formattedValues := strings.Join(fieldValues, ", ")

// 	sqlStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING %s", tableName, formattedColumns, formattedValues, idToReturn)
// 	log.Println(sqlStatement)
// 	var createdEntry interface{}
// 	err := DB.QueryRow(sqlStatement).Scan(&createdEntry)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("New Record is: ", createdEntry)
// 	return createdEntry
// }

// func insertNewData(DB *sql.DB, table string, column string, data string, idToReturn string) int {
// 	cleanedData := cleanData(data)
// 	sqlStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s') RETURNING %s", table, column, cleanedData, idToReturn)
// 	log.Println(sqlStatement)
// 	createdEntryId := 0
// 	err := DB.QueryRow(sqlStatement).Scan(&createdEntryId)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("New Record is: ", createdEntryId, " ", data)
// 	return createdEntryId
// }

// func main() {
// 	args := os.Args[1:]

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	env := &services.Env{DB: db}

// 	var dbConnectErr error
// 	DB, dbConnectErr = sqlx.Connect("postgres", psqlInfo)
// 	if dbConnectErr != nil {
// 		log.Fatalln(dbConnectErr)
// 	}

// 	fmt.Println("DB Connected")

// 	tcgClient := tcg.NewClient(tcgApiKey)

// 	if len(args) == 1 && args[0] == "full" {
// 		fullSyncFlag = true
// 	} else if len(args) == 1 && args[0] == "half" {
// 		halfSyncFlag = true
// 	}

// 	if halfSyncFlag {
// 		sets, err := tcgClient.GetSets()
// 		if err != nil {
// 			panic(err)
// 		}
// 		for _, curSet := range sets {
// 			log.Println(curSet)
// 			// Sync Series to DB first
// 			if !(doesDataExist("series", "series_name", curSet.Series)) {
// 				// insertNewData(DB, "series", "series_name", curSet.Series, "series_id")
// 				newSeries := Series{curSet.Series}
// 				insertNewData1(newSeries, "series_name")
// 			}
// 			// Sync Set to DB
// 			if !(doesDataExist("sets", "api_id", curSet.ID)) {
// 				// newSet := Set{
// 				// 	curSet.ID,
// 				// 	curSet.Name,
// 				// }
// 			}
// 		}
// 		log.Println("Half Sync Complete")
// 	}
// 	if fullSyncFlag {
// 		// Sync Types to DB
// 		types, err := tcgClient.GetTypes()
// 		if err != nil {
// 			panic(err)
// 		}
// 		for _, curType := range types {
// 			log.Println(curType)
// 			if !(doesDataExist("types", "type_name", curType)) {
// 				// insertNewData(DB, "types", "type_name", curType, "type_id")
// 			}
// 		}
// 		// Sync Subtypes to DB
// 		subtypes, err := tcgClient.GetSubTypes()
// 		if err != nil {
// 			panic(err)
// 		}
// 		for _, curSubType := range subtypes {
// 			log.Println(curSubType)
// 			if !(doesDataExist("subtypes", "subtype_name", curSubType)) {
// 				// insertNewData(DB, "subtypes", "subtype_name", curSubType, "subtype_id")
// 			}
// 		}
// 		// Sync Supertypes to DB
// 		supertypes, err := tcgClient.GetSuperTypes()
// 		if err != nil {
// 			panic(err)
// 		}
// 		for _, curSuperType := range supertypes {
// 			log.Println(curSuperType)
// 			if !(doesDataExist("supertypes", "supertype_name", curSuperType)) {
// 				// insertNewData(DB, "supertypes", "supertype_name", curSuperType, "supertype_id")
// 			}
// 		}
// 		// Sync Rarities to DB
// 		rarities, err := tcgClient.GetRarities()
// 		if err != nil {
// 			panic(err)
// 		}
// 		for _, curRarity := range rarities {
// 			log.Println(curRarity)
// 			if !(doesDataExist("rarities", "rarity_name", curRarity)) {
// 				// insertNewData(DB, "rarities", "rarity_name", curRarity, "rarity_id")
// 			}
// 		}
// 		log.Println("Full Sync Complete")
// 		// PriceTypes will remain constant
// 		// Series will remain constant
// 	}

// 	// Sync Cards and Price data to DB
// 	log.Println("Sync Complete")
// }

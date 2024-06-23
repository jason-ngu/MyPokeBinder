package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	_ "github.com/lib/pq"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "4583"
	dbname    = "mypokebinder"
	tcgApiKey = "8c07ea97-a973-43ac-92a9-45d27980d6c6"
)

type Series struct {
	series_name string
}

type Sets struct {
	api_id              string
	set_name            string
	series_id           int
	ptcgo_code          string
	card_total          string
	extended_card_total string
	set_release_date    time.Time
	symbol_image        string
	logo_image          string
	sync_date_created   time.Time
	sync_date_updated   time.Time
}

var fullSyncFlag, halfSyncFlag = false, false

func doesDataExist(db *sql.DB, table string, column string, dataLookupValue string) bool {
	cleanedData := cleanData(dataLookupValue)
	sqlStatement := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = '%s'", table, column, cleanedData)
	var foundRows bool
	err := db.QueryRow(sqlStatement).Scan(&foundRows)
	log.Println(foundRows)
	if err != nil {
		panic(err)
	}
	return foundRows
}

func cleanData(data string) string {
	return strings.ReplaceAll(data, "'", "''")
}

func insertNewData(db *sql.DB, table string, column string, data string, idToReturn string) int {
	cleanedData := cleanData(data)
	sqlStatement := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s') RETURNING %s", table, column, cleanedData, idToReturn)
	log.Println(sqlStatement)
	createdEntryId := 0
	err := db.QueryRow(sqlStatement).Scan(&createdEntryId)
	if err != nil {
		panic(err)
	}
	fmt.Println("New Record is: ", createdEntryId, " ", data)
	return createdEntryId
}

func main() {
	args := os.Args[1:]

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB Connected")

	tcgClient := tcg.NewClient(tcgApiKey)

	if len(args) == 1 && args[0] == "full" {
		fullSyncFlag = true
	} else if len(args) == 1 && args[0] == "half" {
		halfSyncFlag = true
	}

	if halfSyncFlag {
		sets, err := tcgClient.GetSets()
		if err != nil {
			panic(err)
		}
		for _, curSet := range sets {
			log.Println(curSet)
			// Sync Series to DB first
			if !(doesDataExist(db, "series", "series_name", curSet.Series)) {
				insertNewData(db, "series", "series_name", curSet.Series, "series_id")
			}
			// Sync Set to DB
			// if !(doesDataExist(db, "sets", "api_id", curSet.ID)) {

			// }
		}
		log.Println("Half Sync Complete")
	}
	if fullSyncFlag {
		// Sync Types to DB
		types, err := tcgClient.GetTypes()
		if err != nil {
			panic(err)
		}
		for _, curType := range types {
			log.Println(curType)
			if !(doesDataExist(db, "types", "type_name", curType)) {
				insertNewData(db, "types", "type_name", curType, "type_id")
			}
		}
		// Sync Subtypes to DB
		subtypes, err := tcgClient.GetSubTypes()
		if err != nil {
			panic(err)
		}
		for _, curSubType := range subtypes {
			log.Println(curSubType)
			if !(doesDataExist(db, "subtypes", "subtype_name", curSubType)) {
				insertNewData(db, "subtypes", "subtype_name", curSubType, "subtype_id")
			}
		}
		// Sync Supertypes to DB
		supertypes, err := tcgClient.GetSuperTypes()
		if err != nil {
			panic(err)
		}
		for _, curSuperType := range supertypes {
			log.Println(curSuperType)
			if !(doesDataExist(db, "supertypes", "supertype_name", curSuperType)) {
				insertNewData(db, "supertypes", "supertype_name", curSuperType, "supertype_id")
			}
		}
		// Sync Rarities to DB
		rarities, err := tcgClient.GetRarities()
		if err != nil {
			panic(err)
		}
		for _, curRarity := range rarities {
			log.Println(curRarity)
			if !(doesDataExist(db, "rarities", "rarity_name", curRarity)) {
				insertNewData(db, "rarities", "rarity_name", curRarity, "rarity_id")
			}
		}
		log.Println("Full Sync Complete")
		// PriceTypes will remain constant
		// Series will remain constant
	}

	// Sync Cards and Price data to DB
	log.Println("Sync Complete")
}

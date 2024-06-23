package sqldb

import (
	"database/sql"
	"fmt"
)

const (
	host      = "localhost"
	port      = 5432
	user      = "postgres"
	password  = "4583"
	dbname    = "mypokebinder"
	tcgApiKey = "8c07ea97-a973-43ac-92a9-45d27980d6c6"
)

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
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

	return db
}

package handlers

import (
	"backend/common"
	internal "backend/internal"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type handler struct {
	ENV    *internal.Env
	Config common.Configuration
}

func New() handler {
	config := common.SetupConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DatabaseName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	env := &internal.Env{DB: db}
	return handler{ENV: env, Config: config}
}

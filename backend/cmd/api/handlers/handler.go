package handlers

import (
	"backend/config"
	"backend/pkg/db"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type handler struct {
	Config *config.Config
	DB     *sqlx.DB
}

func New() *handler {
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

	return &handler{
		Config: config,
		DB:     db,
	}
}

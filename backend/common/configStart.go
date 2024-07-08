package common

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	ApiKey   string
	Database struct {
		Host         string
		Port         int
		User         string
		Password     string
		DatabaseName string
	}
}

func SetupConfig() Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

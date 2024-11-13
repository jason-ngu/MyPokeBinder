package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	TCGApiKey string
	Database  struct {
		Host         string
		Port         int
		User         string
		Password     string
		DatabaseName string
	}
	Oauth struct {
		Google struct {
			ClientID     string
			ClientSecret string
		}
	}
}

func OpenConfig(configPath string) (*os.File, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ParseConfig(file *os.File) (*Config, error) {
	decoder := json.NewDecoder(file)
	defer file.Close()
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

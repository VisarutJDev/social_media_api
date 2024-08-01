package config

import (
	"encoding/json"
	"log"
	"os"
)

var Config Configuration

type Configuration struct {
	JwtKey   string `json:"jwtKey"`
	MongoURI string `json:"mongoURI"`
	Database string `json:"database"`
}

func LoadConfig(configPath string) {
	configFile, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer configFile.Close()
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&Config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}
}

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	SheetId   string `json:"sheet_id"`
	Columns   string `json:"colunms"`
	TableName string `json:"table"`
}

func ParseConfig() Configuration {
	confPath, err := filepath.Abs("./config/table_config.json")
	if err != nil {
		log.Fatalf("Unable make absolute path: %v", err)
	}
	file, err := os.Open(confPath)
	if err != nil {
		log.Fatalf("Unable open file: %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

type Credentials struct {
	Email        string `json:"client_email"`
	PrivateKey   string `json:"private_key"`
	PrivateKeyID string `json:"private_key_id"`
	TokenURL     string `json:"https://oauth2.googleapis.com/token"`
}

func ParseCredentials() Credentials {
	credPath, err := filepath.Abs("./config/credentials.json")
	if err != nil {
		log.Fatalf("Unable make absolute path: %v", err)
	}
	file, err := os.Open(credPath)
	if err != nil {
		log.Fatalf("Unable open file: %v", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	credentials := Credentials{}
	err = decoder.Decode(&credentials)
	if err != nil {
		fmt.Println("error:", err)
	}
	return credentials
}

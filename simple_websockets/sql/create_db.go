package create_db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func loadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		log.Printf("Error opening config file: %v", err)
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Printf("Error parsing config file: %v", err)
		return config, err
	}
	return config, err
}

func SetupDatabase() (*sql.DB, error) {
	var filePath string = "/home/andreas/Documents/database_access/postgres_config.json"
	config, err := loadConfig(filePath)
	if err != nil {
		log.Println("Error loading config file")
		return nil, err
	}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
	log.Printf("Hallo1")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	log.Printf("Hallo2")
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS messages (
        id SERIAL PRIMARY KEY,
        sender TEXT NOT NULL,
        timestamp TEXT NOT NULL,
        encrypted_message TEXT NOT NULL
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}
	log.Printf("Hallo3")
	return db, nil
}

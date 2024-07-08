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

type Message struct {
	chatId       string
	sender       string
	text         string
	timestamp_ms int64
	hash         string
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

func CreateMessagesTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS messages (
        id SERIAL PRIMARY KEY,
        chatId TEXT NOT NULL,
        sender TEXT NOT NULL,
        text TEXT NOT NULL,
        timestamp_ms BIGINT NOT NULL,
        hash TEXT NOT NULL,
        receivedByClients BOOLEAN DEFAULT FALSE
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

//func WriteMessage(db *sql.DB, chatId string, text string, sender string) error {}

func ConnectToDatabase() (*sql.DB, error) {
	var filePath string = "/home/andreas/Documents/database_access/postgres_config.json"
	config, err := loadConfig(filePath)
	if err != nil {
		log.Println("Error loading config file")
		return nil, err
	}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {

		return nil, err
	}
	return db, nil
}

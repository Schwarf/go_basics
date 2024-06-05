package create_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "bambi"
	password = "SchomerSchabatt"
	dbname   = "chat_server"
)

func SetupDatabase() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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

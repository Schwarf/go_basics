package create_db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func SetupDatabase() (*sql.DB, error) {
	connectionString := "user=username dbname=yourdbname sslmode=disable password=yourpassword"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS messages (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        sender TEXT NOT NULL,
        timestamp TEXT NOT NULL,
        encrypted_message TEXT NOT NULL
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

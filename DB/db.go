package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {

	var err error
	DB, err = sql.Open("sqlite3", "event_api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id            INTEGER PRIMARY KEY AUTOINCREMENT,
		title         TEXT NOT NULL,
		description   TEXT NOT NULL,
		location      TEXT NOT NULL,
		date_time     DATETIME NOT NULL,
		user_id       INTEGER,
		banner_image  TEXT NOT NULL,
		category      TEXT NOT NULL,
		created_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create events table: %v", err)
	}
}

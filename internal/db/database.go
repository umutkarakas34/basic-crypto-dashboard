package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var Database *sql.DB

func InitDB(filepath string) {
	var err error
	Database, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = Database.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection successfully established!")
}

func CloseDB() {
	if err := Database.Close(); err != nil {
		log.Printf("Failed to close database: %v", err)
	} else {
		log.Println("Database connection closed successfully!")
	}
}

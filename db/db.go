package db

import (
	"database/sql"
	"log"
)

// NewMySQLStorageFromURL creates a DB connection using a full DSN
func NewMySQLStorageFromURL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping DB:", err)
		return nil, err
	}

	log.Println("✅ Connected to MySQL via DB_URL!")
	return db, nil
}

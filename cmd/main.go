package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yordanos-habtamu/EcomGo.git/cmd/api"
	"github.com/yordanos-habtamu/EcomGo.git/db"
)

func main() {
	// Read full MySQL URL from env
	dsn := os.Getenv("MYSQL_PUBLIC_URL")
	if dsn == "" {
		log.Fatal("DB_URL environment variable is missing")
	}

	// Connect to MySQL using the URL
	database, err := db.NewMySQLStorageFromURL(dsn)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(database)

	// Start API server
	server := api.NewApiServer(":"+os.Getenv("PORT"), database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("✅ Database connected successfully")
}

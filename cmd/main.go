package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/yordanos-habtamu/EcomGo.git/cmd/api"
	"github.com/yordanos-habtamu/EcomGo.git/db"
)

func main() {
	// Build MySQL config using Railway variables or local fallbacks
	cfg := mysql.Config{
		User:                 getEnv("RAILWAY_MYSQL_USER", "root"),
		Passwd:               getEnv("RAILWAY_MYSQL_PASSWORD", "password"),
		Addr:                 fmt.Sprintf("%s:%s",
			getEnv("RAILWAY_MYSQL_HOST", "127.0.0.1"),
			getEnv("RAILWAY_MYSQL_PORT", "3306")),
		DBName:               getEnv("RAILWAY_MYSQL_DATABASE", "EcomGo"),
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Connect to the database
	database, err := db.NewMysqlStorage(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ping database to ensure connection is alive
	if err := database.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}
	log.Println("Database connected successfully")

	// Start the API server
	server := api.NewApiServer(":"+getEnv("PORT", "8080"), database)
	if err := server.Run(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// getEnv reads an environment variable or returns fallback
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

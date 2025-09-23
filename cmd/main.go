package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/yordanos-habtamu/EcomGo.git/cmd/api"
	"github.com/yordanos-habtamu/EcomGo.git/db"
)

func main() {
	// Read Railway env vars directly
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PWD"),
		Addr:                 os.Getenv("PUBLIC_HOST") + ":" + os.Getenv("MYSQLPORT"),
		DBName:               os.Getenv("DB_NAME"),
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	database, err := db.NewMysqlStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(database)

	server := api.NewApiServer(":"+os.Getenv("PORT"), database)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected successfully")
}

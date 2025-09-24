package main

import (
    "database/sql"
    "log"
    "net/url"
    "os"
    "strings"

    _ "github.com/go-sql-driver/mysql"
    "github.com/yordanos-habtamu/EcomGo.git/cmd/api"
    "github.com/yordanos-habtamu/EcomGo.git/db"
)

func main() {
    // Read Railway DSN from env
    rawDSN := os.Getenv("MYSQL_PUBLIC_URL")
    if rawDSN == "" {
        log.Fatal("MYSQL_PUBLIC_URL environment variable is missing")
    }

    // Convert mysql://... to Go MySQL driver format
    parsed, err := url.Parse(rawDSN)
    if err != nil {
        log.Fatalf("Invalid DSN: %v", err)
    }
    user := parsed.User.Username()
    pass, _ := parsed.User.Password()
    host := parsed.Host
    dbname := strings.TrimPrefix(parsed.Path, "/")
    goDSN := user + ":" + pass + "@tcp(" + host + ")/" + dbname + "?parseTime=true"

    // Connect to MySQL using the converted DSN
    database, err := db.NewMySQLStorageFromURL(goDSN)
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
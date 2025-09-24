package main

import (
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	"github.com/yordanos-habtamu/EcomGo.git/db"
)

func main(){
	
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
		db, err := db.NewMySQLStorageFromURL(goDSN)
    if err != nil {
        log.Fatal(err)
    }


	driver,err := mysql.WithInstance(db,&mysql.Config{})
	if err != nil{
		log.Fatal(err)
	}
  
	m,err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil{
		log.Fatal(err)
	}
	cmd := os.Args[(len(os.Args)-1)]
	if cmd == "up"{
         if err := m.Up(); err != nil && err != migrate.ErrNoChange{
			log.Fatal(err)
		 }
	}
	if cmd == "down"{
		if err := m.Down(); err != nil && err != migrate.ErrNoChange{
			log.Fatal(err)
		 }
	}
}
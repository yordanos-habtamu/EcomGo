package main

import (
	"log"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"os"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/yordanos-habtamu/EcomGo.git/db"
	"github.com/golang-migrate/migrate/v4"
)

func main(){
	
	dsn := os.Getenv("MYSQL_PUBLIC_URL")
	if dsn == "" {
		log.Fatal("MYSQL_PUBLIC_URL environment variable is missing")
	}

	db, err := db.NewMySQLStorageFromURL(dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil{
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
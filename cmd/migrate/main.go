package main

import (
	"log"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"os"
	_ "github.com/golang-migrate/migrate/v4/source/file"
mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/yordanos-habtamu/EcomGo.git/config"
	"github.com/yordanos-habtamu/EcomGo.git/db"
	"github.com/golang-migrate/migrate/v4"
)

func main(){
	db,err := db.NewMysqlStorage(mysqlCfg.Config{
		User:config.Envs.DB_USER,
		Passwd: config.Envs.DB_PWD,
		Addr: config.Envs.DB_ADDR,
		DBName: config.Envs.DB_NAME,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
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
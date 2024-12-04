package main

import (
	"database/sql/driver"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/yordanos-habtamu/EcomGo.git/config"
	"github.com/yordanos-habtamu/EcomGo.git/db"
)

func main(){
	db,err := db.NewMysqlStorage(mysql.Config{
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
	driver,err:= mysql.WithInstance(db, &mysql.Config{})
	if err != nil{
		log.Fatal(err)
	}
	m,err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver
	)
}
package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/yordanos-habtamu/EcomGo.git/cmd/api"
	"github.com/yordanos-habtamu/EcomGo.git/config"
	"github.com/yordanos-habtamu/EcomGo.git/db"
)

func main() {
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
    
	initStorage(db)

	server := api.NewApiServer(":8080",db)
   if error:= server.Run(); error!=nil{
	log.Fatal(error)
   }
}

func initStorage(db *sql.DB){
	err:=db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	log.Println("Database connected successfully")
}
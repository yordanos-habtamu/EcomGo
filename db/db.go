package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func NewMysqlStorage(cfg mysql.Config) (*sql.DB,error) {
   db,err := sql.Open(
     "mysql",
     cfg.FormatDSN(),
   )
   if err != nil{
    log.Fatal(err)
   }
    return db,err
}
	

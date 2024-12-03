package db




import (
    "database/sql"
    "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

type mysql struct {
    // Define the fields of the mysql struct
}

func NewMysqlStorage(cfg mysql) *sql.DB {
    dataSourceName := "username:password@tcp(localhost:3306)/dbname" // Replace with your actual MySQL connection string
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        panic(err.Error()) // Handle the error appropriately
    }
    return db
}
	

package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDatabase() {
	log.Println("Initiate DB")
	// Initialize the database connection here
	// ...
}

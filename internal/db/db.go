package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	username string = "root"
	password string = "Mumbai@02"
	host     string = "localhost"
	port     string = "3306"
	database string = "TODO"
)

func InitDatabase() {
	log.Println("Initiating DB")

	var err error
	DB, err = connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
}

func connectToDatabase() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("DB connection failed", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("DB connection failed", err)
		return nil, err
	}
	log.Println("DB Connected successfully!!!")
	return db, nil
}

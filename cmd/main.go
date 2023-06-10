package main

import (
	"log"
	// "github.com/dhananjaykumardubey/TodoList/internal/handlers"
)

func main() {
	log.Println("Hello, World!")

	db.initDatabase()

	// var router = handlers.setUpRouter()

	// log.Println("Routes", router)

	// // Start the API server
	// log.Fatal(http.ListenAndServe(":8080", router))
}

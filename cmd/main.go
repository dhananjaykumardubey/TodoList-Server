package main

import (
	"log"

	"github.com/dhananjaykumardubey/TodoList-Server/internal/db"
	"github.com/dhananjaykumardubey/TodoList-Server/internal/routes"
)

func main() {
	log.Println("Hello, World!")

	db.InitDatabase()

	var router = routes.SetUpRouter()

	log.Println("Routes", router)
}

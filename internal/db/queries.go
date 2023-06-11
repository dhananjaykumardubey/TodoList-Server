package db

import (
	"fmt"
	"log"
	"time"

	"github.com/dhananjaykumardubey/TodoList-Server/internal/models"
)

func GetTodo() ([]models.TodoResponseModel, error) {
	row, err := DB.Query("Select * from TodoList")
	if err != nil {
		log.Println("Failed to execute the query", err)
		return nil, err
	}
	defer row.Close()

	todolist := make([]models.TodoResponseModel, 0)

	for row.Next() {
		var todo models.TodoResponseModel
		var createdAt time.Time
		err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &createdAt)
		if err != nil {
			log.Println("Failed to parse the data", err)
			return nil, err
		}
		todo.CreatedAt = createdAt
		todolist = append(todolist, todo)

	}
	fmt.Println("got all the data", todolist)
	return todolist, nil
}

func AddTodo(todo models.TodoRequestModel) (int64, error) {
	// Execute INSERT query
	result, err := DB.Exec("INSERT INTO TodoList (Title, completed) VALUES (?, ?)", todo.Title, false)
	if err != nil {
		log.Println("Failed to execute query:", err)
		return 0, err
	}

	// Get the ID of the inserted row
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Failed to retrieve last insert ID:", err)
		return 0, err
	}

	fmt.Println("New Todo ID:", id)
	return id, nil
}

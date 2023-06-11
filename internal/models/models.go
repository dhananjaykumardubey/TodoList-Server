package models

import (
	"time"
)

type TodoResponseModel struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoRequestModel struct {
	Title string `json:"title"`
}

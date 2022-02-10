package models

import (
	"github.com/google/uuid"
)

type Todo struct {
	Id          uuid.UUID `json:"id"`
	Author      uuid.UUID `json:"todo_author"`
	Title       string    `json:"todo_title"`
	Description string    `json:"todo_description"`
	Done        bool      `json:"done"`
}

type CreateTodo struct {
	Id          uuid.UUID `json:"id"`
	Author      uuid.UUID `json:"todo_author"`
	Title       string    `json:"todo_title"`
	Description string    `json:"todo_description"`
}

type TodosResponseJson struct {
	Success bool   `json:"success"`
	Data    []Todo `json:"todos"`
	Message string `json:"message"`
}

type TodoResponseJson struct {
	Success bool   `json:"success"`
	Data    Todo   `json:"todos"`
	Message string `json:"message"`
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Author      uuid.UUID `json:"todo_author"`
	Title       string    `json:"todo_title"`
	Description string    `json:"todo_description"`
	Done        bool      `json:"done"`
}

type CreateTodo struct {
	Title       string
	Description string
	Done        bool
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

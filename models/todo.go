package models

import (
  "time"
  "github.com/google/uuid"
)

type Todo struct {
  id uuid.UUID `gorm:"primaryKey"json:"_id"`
  author uuid.UUID `json:"todo_author"`
  title string `json:"todo_title"`
  description string `json:"todo_description"`
  done bool `json:"done"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"created_at"`
}

type CreateTodo struct {
  title string
  description string
  done bool
}
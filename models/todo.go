package models

type Todo struct {
  id int32
  title string
  description string
  done bool
}

type NewTodo struct {
  title string
  description string
}
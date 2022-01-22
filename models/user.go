package models

import (
  "net/http"
  "fmt"
  "database/sql"
)

type User struct {
  id int32
  username string
  password string
  email string
  todos []Todo
}

type CreateUser struct {
  username string
  password string
  email string
}

func NewUser(user CreateUser) User {
  return User {
    username: user.username,
    password: user.password,
    email: user.email,
    todos: []Todo{},
  }
}

func (user *User) VerifyPassword(password string) bool {
  if user.password == password {
    return true
  }
  return false
}
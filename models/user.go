package models

import (
  "net/http"
  "fmt"
  "database/sql"
)

type User struct {
  Id int32 `json:"_id"`
  Username string `json:"username"`
  Password string `json:"password"`
  Email string `json:"email"`
  Todos []Todo `json:"todos"`
}

type CreateUser struct {
  Username string
  Password string
  Email string
}

type UserJsonResponse struct {
  Type    string `json:"type"`
  Data    []User `json:"data"`
  Message string `json:"message"`
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
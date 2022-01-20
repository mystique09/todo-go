package user

import (
  "server-go/todo"
  "net/http"
  "fmt"
  "database/sql"
)

type User struct {
  id int32
  username string
  password string
  email string
  todos []todo.Todo
}

type NewUser struct {
  username string
  password string
  email string
}

func New(user NewUser) User {
  return User {
    username: user.username,
    password: user.password,
    email: user.email,
    todos: []todo.Todo{},
  }
}

func (user *User) VerifyPassword(password string) bool {
  if user.password == password {
    return true
  }
  return false
}

// ROUTES

/* Index user route */
func AllUser(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
      http.NotFound(w,r)
      return
    }
    var users, err = db.Query("SELECT * FROM User;")
    if err != nil {
      panic(err)
    }
    fmt.Fprintf(w, "Hello, from user route! %v", *users)
  }
}

/* Create user route */
func CreateNewUser(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      http.NotFound(w,r)
      return
    }
    fmt.Fprintf(w, "Hello, from create new user route!")
  }
}
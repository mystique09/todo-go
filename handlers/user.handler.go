package handlers

import (
  "database/sql"
  "net/http"
  "fmt"
)

func AllUser(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
      http.NotFound(w,r)
      return
    }
    
    fmt.Fprintf(w, "All users route.")
  }
}

/* Create user route */
func CreateNewUser(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      http.NotFound(w,r)
      return
    }
    
    fmt.Fprintf(w, "Create new user route.")
  }
}
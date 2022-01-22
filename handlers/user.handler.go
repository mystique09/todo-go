package handlers

import (
  "database/sql"
  "net/http"
  "server-go/utils"
  "server-go/models"
  "encoding/json"
)

func AllUser(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
      http.NotFound(w, r)
      return
    }
    
    rows, err := db.Query(`SELECT id, username, email FROM "User"`)
    utils.CheckError(err)
    
    var users []models.QueryableUser = []models.QueryableUser{}
    
    for rows.Next() {
      var id int32
      var username string
      var email string
      
      err = rows.Scan(&id, &username, &email)
      utils.CheckError(err)
      users = append(users, models.QueryableUser { Username: username, Email: email})
    }
    var response = models.UserJsonResponse {
      Success: true,
      Data: users,
      Message: "All users.",
    }
    json.NewEncoder(w).Encode(response)
  }
}

/* Create user route */
func CreateNewUser(db *sql.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      http.NotFound(w, r)
      return
    }
    var username string = r.FormValue("username")
    var password string = r.FormValue("password")
    var email string = r.FormValue("email")
    var response utils.Response
    
    if username == "" || password == "" || email == "" {
      response = utils.Response {
        Success: false,
        Message: "Missing required fields!",
      }
      json.NewEncoder(w).Encode(response)
      return
    } else {
      var lastInsertedId int32 
      err := db.QueryRow(`INSERT INTO "User"(username, password, email) VALUES($1, $2, $3) returning id;`, username, password, email).Scan(&lastInsertedId)
      
      utils.CheckError(err)
      
      response = utils.Response {
        Success: true,
        Message: "New user added with id of ",
      }
      json.NewEncoder(w).Encode(response)
      return
    }
  }
}
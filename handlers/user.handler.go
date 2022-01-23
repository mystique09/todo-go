package handlers

import (
  "gorm.io/gorm"
  "net/http"
  "server-go/utils"
  "fmt"
  //"time"
  "server-go/models"
  "encoding/json"
  "github.com/google/uuid"
)

func AllUser(db *gorm.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
      http.NotFound(w, r)
      return
    }
    
    var users []models.User
    db.Find(&users)
    
    var response models.UserJsonResponse = models.UserJsonResponse {
      Success: true,
      Data: users,
      Message: "All users.",
    }
    json.NewEncoder(w).Encode(response)
  }
}

/* Create user route */
func CreateNewUser(db *gorm.DB) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      http.NotFound(w, r)
      return
    }
    
    var username string = r.FormValue("username")
    var password string = r.FormValue("password")
    var email string = r.FormValue("email")
    var response utils.Response
    
    if username == "" || email == "" || password == "" {
      w.WriteHeader(http.StatusBadRequest)
      response = utils.Response {
        Success: false,
        Message: "Missing required fields!",
      }
      json.NewEncoder(w).Encode(response)
      return
    }
    
    var new_uuid uuid.UUID = uuid.New()
    
    var new_user models.User = models.User {
      ID: new_uuid,
      Username: username,
      Password: password,
      Email: email,
    }
    
    if err := db.Create(&new_user).Error; err != nil {
      w.WriteHeader(http.StatusBadRequest)
      response = utils.Response {
        Success: false,
        Message: fmt.Sprintf("%v", err),
      }
      json.NewEncoder(w).Encode(response)
      return
    }
    
    response = utils.Response {
      Success: true,
      Message: "New user added!",
    }
    json.NewEncoder(w).Encode(response)
    }
  }
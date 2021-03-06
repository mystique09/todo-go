package handlers

import (
  "gorm.io/gorm"
  "net/http"
  "server-go/utils"
  //"fmt"
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
    
    var users []models.QueryableUser
    db.Model(&models.User{}).Select("id", "username", "email", "created_at", "updated_at").Find(&users)
    
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

    var new_user models.User 
    utils.ParseBody(r, &new_user)
    new_user.Id = uuid.New()
    
    response := utils.Response {
        Success: false,
        Message: "",
      }
    
    if new_user.Username == "" || new_user.Email == "" || new_user.Password == "" {
        response.Message = "Missing required fields."
        w.WriteHeader(http.StatusBadRequest)
        w.Write(utils.ParseJson(response))
        return
    }
    
    var hasUser models.QueryableUser
    db.Model(&models.User{}).Where("username = ?", new_user.Username).Select("id", "username", "email", "created_at", "updated_at").Find(&hasUser)
    
    if hasUser.Username != "" {
      response.Message = "User already exist!"
      w.WriteHeader(http.StatusBadRequest)
      w.Write(utils.ParseJson(response))
      return
    }
    
    if err := db.Create(&new_user).Error; err != nil {
      response.Message = "Error while creating user.";
      w.WriteHeader(http.StatusBadRequest)
      w.Write(utils.ParseJson(response))
      return
    }
    
    response.Message = "User created successfully."
    response.Success = true
    w.WriteHeader(http.StatusOK)
    w.Write(utils.ParseJson(response))
    }
  }
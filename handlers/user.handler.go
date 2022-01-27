package handlers

import (
	"fmt"
	"net/http"
	"server-go/utils"

	"gorm.io/gorm"

	//"time"
	"encoding/json"
	"server-go/models"

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

		var response models.UserJsonResponse = models.UserJsonResponse{
			Success: true,
			Data:    users,
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

		response := utils.Response{
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
			response.Message = "Error while creating user."
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

func GetUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}
		var user_uuid = r.URL.Query()["uuid"][0]
		var user models.QueryableUser
		var response models.UserJsonResponse = models.UserJsonResponse{
			Success: false,
			Data:    []models.QueryableUser{},
		}

		if user_uuid == "" {
			response.Message = "Missing uuid query parameter."
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ParseJson(response))
			return
		}

		db.Model(&models.User{}).Where("id = ?", user_uuid).Find(&user)

		if user.Username == "" {
			response.Message = fmt.Sprintf("No user found with id %s", user_uuid)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ParseJson(response))
			return
		}
		response.Success = true
		response.Message = "User found."
		response.Data = append(response.Data, user)
		w.WriteHeader(http.StatusOK)
		w.Write(utils.ParseJson(response))
	}
}

func UpdateUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.NotFound(w, r)
			return
		}

		var user_uuid = r.URL.Query()["uuid"][0]
		var prop = r.URL.Query()["prop"][0]
		var hasUser models.QueryableUser
		var user models.User
		response := utils.Response{
			Success: true,
			Message: "",
		}

		utils.ParseBody(r, &user)

		if user_uuid == "" || prop == "" {
			response = utils.Response{
				Success: false,
				Message: "Missing some query parameter.",
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ParseJson(response))
			return
		}

		db.Model(&models.User{}).Where("id = ?", user_uuid).Find(&hasUser)

		if hasUser.Username == "" {
			response.Success = false
			response.Message = fmt.Sprintf("No user found with id %s", user_uuid)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ParseJson(response))
			return
		}

		switch prop {
		case "password":
			if user.Password != "" {
				db.Model(&models.User{}).Where("id = ?", user_uuid).Update("password", user.Password)
				response.Message = "Password changed."
				break
			}
		case "username":
			if user.Username != "" {
				db.Model(&models.User{}).Where("id = ?", user_uuid).Update("username", user.Username)
				response.Message = "Username changed."
				break
			}
		case "email":
			if user.Email != "" {
				db.Model(&models.User{}).Where("id = ?", user_uuid).Update("email", user.Email)
				response.Message = "Email changed."
				break
			}
		default:
			response.Message = "Unknown prop."
		}
		response.Message = "Update success."
		w.WriteHeader(http.StatusOK)
		w.Write(utils.ParseJson(response))
	}
}

func DeleteUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.NotFound(w, r)
			return
		}

		var user_uuid = r.URL.Query()["uuid"][0]
		var hasUser models.QueryableUser
		var response utils.Response = utils.Response{
			Success: false,
			Message: "",
		}

		if user_uuid == "" {
			response.Message = "Missing uuid query parameter."
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ParseJson(response))
			return
		}
		db.Model(&models.User{}).Where("id = ?", user_uuid).Find(&hasUser)

		if hasUser.Username == "" {
			response.Message = fmt.Sprintf("No user found with id %s", user_uuid)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ParseJson(response))
			return
		}

		db.Where("id = ?", user_uuid).Delete(&models.User{})
		response = utils.Response{
			Success: true,
			Message: "User deleted.",
		}

		w.WriteHeader(http.StatusOK)
		w.Write(utils.ParseJson(response))
	}
}

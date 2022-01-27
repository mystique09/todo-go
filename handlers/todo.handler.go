package handlers

import (
	"net/http"
	"server-go/models"
	"server-go/utils"

	"gorm.io/gorm"
)

func AllTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		var todos []models.Todo

		db.Model(&models.User{}).Find(&todos)

		var response models.TodosResponseJson = models.TodosResponseJson{
			Success: true,
			Data:    todos,
			Message: "All todos",
		}

		w.WriteHeader(http.StatusOK)
		w.Write(utils.ParseJson(response))
	}
}

func GetTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		var todo_uuid = r.URL.Query()["uuid"][0]
		response := models.TodoResponseJson{
			Message: "",
			Success: false,
		}

		if todo_uuid == "" {
			response.Message = "Missing todo uuid."
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ParseJson(response))
			return
		}

		var todo models.Todo
		db.Model(&models.Todo{}).Where("id = ?", todo_uuid).Find(&todo)

		response = models.TodoResponseJson{
			Success: true,
			Data:    todo,
			Message: "Todo",
		}
		w.WriteHeader(http.StatusOK)
		w.Write(utils.ParseJson(response))
	}
}

func CreateTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func UpdateTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

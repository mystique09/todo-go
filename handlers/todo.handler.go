package handlers

import (
	"net/http"
	"server-go/models"
	"server-go/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AllTodo(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		var todos []models.Todo

		db.Model(&models.Todo{}).Find(&todos)

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
		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}
		var n_todo models.Todo
		utils.ParseBody(r, &n_todo)
		n_todo.Id = uuid.New()
		n_todo.Done = false

		var response utils.Response

		if n_todo.Author == uuid.Nil {
			response = utils.Response{
				Message: n_todo.Author.String(),
				Success: false,
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.ParseJson(response))
			return
		}

		if err := db.Create(&n_todo).Error; err != nil {
			response := utils.Response{
				Success: false,
				Message: "Error while creating new todo.",
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(utils.ParseJson(response))
			return
		}
		response = utils.Response{
			Success: true,
			Message: "New todo added.",
		}
		w.WriteHeader(http.StatusOK)
		w.Write(utils.ParseJson(response))
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

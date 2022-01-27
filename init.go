package main

import (
	"log"
	"net/http"
	database "server-go/db"
	"server-go/handlers"
	"server-go/models"

	"gorm.io/gorm"
)

type App struct {
	Mux *http.ServeMux
}

func (app *App) Init() {
	var db *gorm.DB = database.InitDb()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})

	app.Mux = http.NewServeMux()

	// USER HANDLERS
	app.Mux.HandleFunc("/", handlers.IndexRoute)
	app.Mux.HandleFunc("/users", handlers.AllUser(db))
	app.Mux.HandleFunc("/users/create", handlers.CreateNewUser(db))
	app.Mux.HandleFunc("/users/get", handlers.GetUser(db))
	app.Mux.HandleFunc("/users/delete", handlers.DeleteUser(db))
	app.Mux.HandleFunc("/users/update", handlers.UpdateUser(db))

	// TODO HANDLERS
	app.Mux.HandleFunc("/users", handlers.AllTodo(db))
	app.Mux.HandleFunc("/users/create", handlers.CreateTodo(db))
	app.Mux.HandleFunc("/users/get", handlers.GetTodo(db))
	app.Mux.HandleFunc("/users/delete", handlers.DeleteTodo(db))
	app.Mux.HandleFunc("/users/update", handlers.UpdateTodo(db))
}

func (app *App) Run() {
	log.Println("Server started.")
	log.Fatal(http.ListenAndServe(":3000", app.Mux))
}

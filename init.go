package main

import (
	"database/sql"
	"log"
	"net/http"
	database "server-go/db"
	"server-go/handlers"
)

type App struct {
	Mux *http.ServeMux
}

func (app *App) Init() {
	var db *sql.DB = database.InitDb()

	app.Mux = http.NewServeMux()
	app.Mux.HandleFunc("/", handlers.IndexRoute)
	app.Mux.HandleFunc("/users", handlers.AllUser(db))
	app.Mux.HandleFunc("/users/create", handlers.CreateNewUser(db))
}

func (app *App) Run() {
	log.Println("🎆 Server started.")
	log.Fatal(http.ListenAndServe(":3000", app.Mux))
}

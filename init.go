package main

import (
  "net/http"
  "log"
  "server-go/handlers"
  database "server-go/db"
  "database/sql"
)

type App struct {
  Mux *http.ServeMux
}

func (app *App) Init() {
  var db *sql.DB = database.InitDb()
  
  app.Mux = http.NewServeMux()
  app.Mux.HandleFunc("/", handlers.IndexRoute)
  app.Mux.HandleFunc("/user", handlers.AllUser(db))
  app.Mux.HandleFunc("/user/create", handlers.CreateNewUser(db))
}

func (app *App) Run() {
  log.Println("🎆 Server started.")
  log.Fatal(http.ListenAndServe(":3000", app.Mux))
}

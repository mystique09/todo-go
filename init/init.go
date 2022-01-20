package init

import (
  "fmt"
  "net/http"
  "time"
  "log"
  "server-go/user"
  database "server-go/db"
  "database/sql"
  //"server-go/todo"
)

type App struct {
  Mux *http.ServeMux
}

func (app *App) Init() {
  var db *sql.DB = database.InitDb()
  
  app.Mux = http.NewServeMux()
  app.Mux.HandleFunc("/", IndexRoute)
  
  app.Mux.HandleFunc("/user", user.AllUser(db))
  app.Mux.HandleFunc("/user/create", user.CreateNewUser(db))
}

func (app *App) Run() {
  fmt.Println("Server started at", time.Now())
  log.Fatal(http.ListenAndServe(":3000", app.Mux))
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  fmt.Fprintf(w, "Hello, World!")
}
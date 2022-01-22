package handlers

import (
  "fmt"
  "net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  fmt.Fprintf(w, "Hello, World!")
}
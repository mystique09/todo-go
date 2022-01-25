package utils

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	//"errors"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ParseBody(r *http.Request, x interface{}) {
  if body, err := ioutil.ReadAll(r.Body); err == nil {
    if err := json.Unmarshal([]byte(body), x); err != nil {
      return
    }
  }
}

func ParseJson(data interface {}) []byte {
  json, err := json.Marshal(data)
  if err != nil {
    panic(err)
  }
  return json
}

type Response struct {
  Success bool `json:"success"`
  Message string `json:"message"`
}

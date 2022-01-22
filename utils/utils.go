package utils

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Response struct {
  Success bool `json:"success"`
  Message string `json:"message"`
}

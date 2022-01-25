package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"errors"
)

func InitDb() *gorm.DB {
  if err := godotenv.Load(); err != nil{
    log.Fatal("Error loading .env file.")
  }
  
  DB_NAME := os.Getenv("DB_NAME")
  DB_USERNAME := os.Getenv("DB_USERNAME")
  DB_PASSWORD := os.Getenv("DB_PASSWORD")
  DB_HOST := os.Getenv("DB_HOST")
  
  var DB_CONFIG string = fmt.Sprintf("postgres://%s:%s@%s/%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_NAME)
  
	conn, err := gorm.Open(postgres.Open(DB_CONFIG), &gorm.Config{})
	if err != nil {
		log.Fatal(errors.New("Failed to connect to the Database."))
	}
	
	fmt.Println("🎆 Database connected!")
	return conn
}
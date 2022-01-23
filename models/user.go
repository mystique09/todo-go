package models

import (
  "gorm.io/gorm"
  "time"
  "github.com/google/uuid"
)

type User struct {
  gorm.Model
  ID uuid.UUID `gorm:"primaryKey"json:"_id"`
  Username string `gorm:"unique;not null"json:"username"`
  Password string `json:"password"`
  Email string `gorm:"unique;not null"json:"email"`
  /*CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"created_at"`
  DeletedAt time.Time `json:"deleted_at"`*/
}

type QueryableUser struct {
  Id int32 `json:"_id"`
  Username string `json:"username"`
  Email string `json:"email"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"created_at"`
}

type CreateUser struct {
  Username string
  Password string
  Email string
}

type UserJsonResponse struct {
  Success    bool `json:"success"`
  Data    []User `json:"data"`
  Message string `json:"message"`
}
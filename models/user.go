package models

import (
  //"gorm.io/gorm"
  "time"
  "github.com/google/uuid"
)

type User struct {
  Id uuid.UUID `gorm:"primarykey"json:"_id"`
  Username string `gorm:"unique;not null"json:"username"`
  Password string `json:"password"`
  Email string `gorm:"not null"json:"email"`
  CreatedAt time.Time `gorm:"autoCreateTime"json:"created_at"`
  UpdatedAt time.Time `gorm:"autoUpdateTime"json:"updated_at"`
  DeledAt time.Time `gorm:"index"json:"deleted_at"`
}

type QueryableUser struct {
  Id uuid.UUID `json:"_id"`
  Username string `json:"username"`
  Email string `json:"email"`
  CreatedAt time.Time `json:"updated_at"`
  UpdatedAt time.Time `json:"created_at"`
}

type CreateUser struct {
  username string
  password string
  email string
}

type UserJsonResponse struct {
  Success    bool `json:"success"`
  Data    []QueryableUser `json:"data"`
  Message string `json:"message"`
}
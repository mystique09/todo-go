package models

import (
	//"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"_id"`
	Username  string    `gjson:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeledAt   time.Time `json:"deleted_at"`
}

type QueryableUser struct {
	Id        uuid.UUID `json:"_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"updated_at"`
	UpdatedAt time.Time `json:"created_at"`
}

type CreateUser struct {
	Username string
	Password string
	Email    string
}

type UserJsonResponse struct {
	Success bool            `json:"success"`
	Data    []QueryableUser `json:"data"`
	Message string          `json:"message"`
}

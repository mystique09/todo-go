package models

type User struct {
  Id int32 `json:"_id"`
  Username string `json:"username"`
  Password string `json:"password"`
  Email string `json:"email"`
}

type QueryableUser struct {
  Id int32 `json:"_id"`
  Username string `json:"username"`
  Email string `json:"email"`
}

type CreateUser struct {
  Username string
  Password string
  Email string
}

type UserJsonResponse struct {
  Success    bool `json:"success"`
  Data    []QueryableUser `json:"data"`
  Message string `json:"message"`
}

func (user *User) VerifyPassword(password string) bool {
  if user.Password == password {
    return true
  }
  return false
}
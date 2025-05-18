package schema

import (
	"time"
)

type UserCreateRequest struct {
	Firstname  string    `json:"first_name" binding:"required,max=20"`
	Lastname   string    `json:"last_name" binding:"max=20"`
	Email      string    `json:"email" binding:"required,email"`
	Password   string    `json:"password" binding:"required"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

package models

import (
	"time"
)

type User struct {
	Id	int
	FirstName 	string	`json:"first_name" validates:"required"`
	LastName 	string `json:"last_name" validates:"required"`
	Email	string	`json:"email" validates:"email, required"`
	Uid 	string 	`json:"uid"`
	Token 	string	`json:"token"`
	RefreshToken	string	`json:"refresh_token"`
	CreatedAt time.Time 	`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
}

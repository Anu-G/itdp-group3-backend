package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

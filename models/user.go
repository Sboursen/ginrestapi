package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Email   int    `json:"email"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

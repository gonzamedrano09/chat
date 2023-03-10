package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name        string    `json:"name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
}

func (User) TableName() string {
	return "users"
}

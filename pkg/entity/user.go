package entity

import (
	"time"
)

var Genders = [3]string{
	"Male",
	"Female",
	"Undefined",
}

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"index"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

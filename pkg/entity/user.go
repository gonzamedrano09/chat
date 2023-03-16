package entity

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Username    string    `json:"username" gorm:"not null;index;unique"`
	Password    string    `json:"password" gorm:"not null;"`
	FirstName   string    `json:"first_name" gorm:"default:null"`
	LastName    string    `json:"last_name" gorm:"default:null"`
	Gender      string    `json:"gender" gorm:"default:null"`
	DateOfBirth time.Time `json:"date_of_birth"  gorm:"default:null"`
	CreatedAt   time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

var Genders = [3]string{
	"Male",
	"Female",
	"Undefined",
}

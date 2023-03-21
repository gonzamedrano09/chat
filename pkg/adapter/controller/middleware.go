package controller

import "gorm.io/gorm"

type MiddlewareDatabaseConnection struct {
	Database *gorm.DB
}

func NewMiddlewareDatabaseConnection(database *gorm.DB) MiddlewareDatabaseConnection {
	return MiddlewareDatabaseConnection{Database: database}
}

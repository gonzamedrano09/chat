package infrastructure

import (
	"fmt"
	"github.com/gonzamedrano09/chat/pkg/config"
	"github.com/gonzamedrano09/chat/pkg/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDatabase() *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DatabaseUser,
		config.Config.DatabasePassword,
		config.Config.DatabaseHost,
		config.Config.DatabasePort,
		config.Config.DatabaseName)
	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	err = database.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalln(err)
	}
	return database
}

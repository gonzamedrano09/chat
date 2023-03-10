package infrastructure

import (
	"fmt"
	"github.com/gonzamedrano09/chat/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return database
}

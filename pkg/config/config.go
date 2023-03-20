package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Settings struct {
	SecretKey        string `env:"SECRET_KEY"`
	DatabaseUser     string `env:"MYSQL_USER"`
	DatabasePassword string `env:"MYSQL_PASSWORD"`
	DatabaseHost     string `env:"MYSQL_HOST"`
	DatabasePort     string `env:"MYSQL_PORT"`
	DatabaseName     string `env:"MYSQL_NAME"`
}

var Config = Settings{}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading environment variables from .env file")
		fmt.Println(err)
	}
	if err := env.Parse(&Config); err != nil {
		fmt.Println("Error parsing environment variables to Config struct")
		fmt.Println(err)
	}
}

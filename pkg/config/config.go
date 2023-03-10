package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type Settings struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
}

var Config Settings

func LoadConfig() {
	InternalConfig := &Config

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&InternalConfig); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_HOST        string
	DB_PORT        string
	DB_NAME        string
	DB_NAME_TEST   string
	JWT_SECRET_KEY string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{
		DB_USERNAME:    os.Getenv("DB_USERNAME"),
		DB_PASSWORD:    os.Getenv("DB_PASSWORD"),
		DB_HOST:        os.Getenv("DB_HOST"),
		DB_PORT:        os.Getenv("DB_PORT"),
		DB_NAME:        os.Getenv("DB_NAME"),
		DB_NAME_TEST:   os.Getenv("DB_NAME_TEST"),
		JWT_SECRET_KEY: os.Getenv("JWT_SECRET_KEY"),
	}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	} else {
		_ = viper.Unmarshal(cfg)
	}

	Cfg = cfg
}

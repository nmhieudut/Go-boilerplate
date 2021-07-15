package config

import (
	"rest-api/pkg/utils"
)

type Config struct {
	DbUrl      string
	DbName     string
	ServerPort string
	SecrectJWT string
}

var Setting *Config

func init() {
	Setting = &Config{
		DbUrl:      utils.GetEnv("DB_URL"),
		DbName:     utils.GetEnv("DB_NAME"),
		ServerPort: utils.GetEnv("PORT"),
		SecrectJWT: utils.GetEnv("SECRET_JWT"),
	}
}

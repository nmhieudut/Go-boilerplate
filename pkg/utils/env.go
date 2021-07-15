package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading .env file")
	}
	return os.Getenv(key)
}
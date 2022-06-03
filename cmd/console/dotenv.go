package console

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

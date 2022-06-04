package console

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("DOTENV: No .env file found or error reading file.")
	} else {
		log.Printf("DOTENV: .env found.")
	}
}

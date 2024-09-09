package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvLoad(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

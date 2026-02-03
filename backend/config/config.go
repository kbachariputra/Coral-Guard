package config

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)

func loadConfig(key string) string{
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
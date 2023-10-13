package config 

import (
	"github.com/joho/godotenv"
	"os"
)

func EnvVariable(key string) string{

	environment := os.Getenv("GOENV")

	if environment != "PROD" {
		err := godotenv.Load(".env.dev")

		if err != nil {
			panic("Error Loading Environment Variables from file .env.dev")
		}

	}

	if environment == "PROD" {
		err := godotenv.Load(".env.prod")

		if err != nil {
			panic("Error Loading Environment Variables from file .env.prod")
		}
	}

	return os.Getenv(key)
}

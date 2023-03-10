package config 

import (
	"github.com/joho/godotenv"
	"os"
)

func EnvVariable(key string) string{

	err := godotenv.Load(".env")

	if err != nil {
		panic("Error Loading Environment Variables from file")
	}

	return os.Getenv(key)
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := os.Getenv("APP_ENV")

	var file string

	switch env {
	case "prod":
		file = ".env.prod"
	case "stag":
		file = ".env.stag"
	default:
		file = ".env"
	}

	err := godotenv.Overload(file)
	if err != nil {
		log.Fatal("No env file found")
	}
}

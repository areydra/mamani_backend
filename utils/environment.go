package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(types string) {
	var err error

	if types == "production" {
		err = godotenv.Load(".env.production")
	} else {
		err = godotenv.Load(".env.staging")
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

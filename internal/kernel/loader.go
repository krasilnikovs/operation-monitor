package kernel

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	loadDotEnv()
}

func loadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("error during loading .env: %w", err)
	}

	godotenv.Overload(".env.local")
}

package initialisers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env: %v", err)
	}
}

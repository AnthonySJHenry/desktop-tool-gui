package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() error {
	_ = godotenv.Load()
	return nil
}

func Port() string {
	return os.Getenv("PORT")
}

func DatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

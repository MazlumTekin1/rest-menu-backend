package config

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

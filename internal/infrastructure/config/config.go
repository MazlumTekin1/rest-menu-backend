package config

//PATH: internal/infrastructure/config/config.go

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	JWT JWTConfig
}

type JWTConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {

	absPath, err := filepath.Abs("./.env")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	err = godotenv.Load(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %v", err)
	}

	return &Config{
		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}, nil
}

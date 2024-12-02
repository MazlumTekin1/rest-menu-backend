package config

//PATH: internal/infrastructure/config/database.go

import (
	"fmt"
	"rest-menu-service/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		Port:     getEnvOrDefault("DB_PORT", "5432"),
		User:     getEnvOrDefault("DB_USER", "postgres"),
		Password: getEnvOrDefault("DB_PASSWORD", "12345"),
		DBName:   getEnvOrDefault("DB_NAME", "postgres"),
	}
}

func GetConnectionString() string {
	config := NewDatabaseConfig()
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}

func SetupGormDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(GetConnectionString()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&domain.Menu{},
		&domain.Category{},
		&domain.Product{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return db, nil
}

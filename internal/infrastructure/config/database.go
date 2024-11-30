package config

import (
	"fmt"
	"os"
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

func (c *DatabaseConfig) GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// func SetupDatabase() (*pgxpool.Pool, error) {
// 	config := NewDatabaseConfig()
// 	connString := config.GetConnectionString()

// 	poolConfig, err := pgxpool.ParseConfig(connString)
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing connection string: %v", err)
// 	}

// 	// Configure pool settings
// 	poolConfig.MaxConns = 10
// 	poolConfig.MinConns = 2

// 	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating connection pool: %v", err)
// 	}

// 	// Test the connection
// 	if err := pool.Ping(context.Background()); err != nil {
// 		return nil, fmt.Errorf("error connecting to the database: %v", err)
// 	}

// 	return pool, nil
// }

func SetupGormDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&domain.Menu{},
		&domain.Category{},
		&domain.Product{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

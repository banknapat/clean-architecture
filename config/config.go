package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() (*Config, error) {
	once.Do(func() {
		_ = godotenv.Load() // หากมีไฟล์ .env

		cfg = &Config{
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBUser:     getEnv("DB_USER", "postgres"),
			DBPassword: getEnv("DB_PASSWORD", "postgres"),
			DBName:     getEnv("DB_NAME", "mydb"),
			JWTSecret:  getEnv("JWT_SECRET", "SUPER_SECRET_KEY"),
		}
	})

	return cfg, nil
}

func (c *Config) GetPostgresDSN() string {
	// postgres://user:password@host:port/dbname
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

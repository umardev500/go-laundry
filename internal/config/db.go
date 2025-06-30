package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	User    string
	Pass    string
	Host    string
	Name    string
	SSLMode string
	Port    string // Optional, default 5432
}

func LoadDatabaseConfig() *DatabaseConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	return &DatabaseConfig{
		User:    getEnv("DB_USER", "root"),
		Pass:    getEnv("DB_PASS", "root"),
		Host:    getEnv("DB_HOST", "127.0.0.1"),
		Name:    getEnv("DB_NAME", "pos"),
		SSLMode: getEnv("DB_SSLMODE", "disable"),
		Port:    getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// JwtConfig holds JWT-related configuration.
type JwtConfig struct {
	Secret     string
	Expiration time.Duration
	Issuer     string
}

// DBConfig holds DB-related configuration.
type DBConfig struct {
	User    string
	Pass    string
	Host    string
	Name    string
	SSLMode string
}

// AppConfig holds all application-level configuration.
type AppConfig struct {
	Port        int
	Jwt         JwtConfig
	DB          DBConfig
	FiberConfig fiber.Config
}

// Load loads environment variables from .env and builds AppConfig.
func Load() *AppConfig {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, falling back to OS environment")
	}

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid PORT: %v", err)
	}

	expHoursStr := os.Getenv("JWT_EXPIRATION_HOURS")
	expHours, err := strconv.Atoi(expHoursStr)
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRATION_HOURS: %v", err)
	}

	return &AppConfig{
		Port: port,
		Jwt: JwtConfig{
			Secret:     os.Getenv("JWT_SECRET"),
			Expiration: time.Duration(expHours) * time.Hour,
			Issuer:     os.Getenv("JWT_ISSUER"),
		},
		DB: DBConfig{
			User:    os.Getenv("DB_USER"),
			Pass:    os.Getenv("DB_PASS"),
			Host:    os.Getenv("DB_HOST"),
			Name:    os.Getenv("DB_NAME"),
			SSLMode: os.Getenv("DB_SSLMODE"),
		},
		FiberConfig: fiber.Config{
			DisableStartupMessage: true,
		},
	}
}

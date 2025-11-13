package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	Port int
}

type DatabaseConfig struct {
	Url string
}

type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

func LoadConfig() *Config {
	cfg := &Config{}
	// App config
	cfg.App.Port = getEnvAsInt("APP_PORT", 8080)

	// Database config
	cfg.Database.Url = getEnv("DATABASE_URL", "")

	return cfg
}

// Helper to read env variable or fallback to default
func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return defaultVal
}

// Helper to read int env variable or fallback to default
func getEnvAsInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}

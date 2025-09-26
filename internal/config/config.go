package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Dsn    string `mapstructure:"dsn"`
}

type Firebase struct {
	CredentialsFile string `mapstructure:"credentials_file"`
}

type JWT struct {
	Issuer                    string `mapstructure:"issuer"`
	Secret                    string `mapstructure:"secret"`
	ExpirySeconds             int64  `mapstructure:"expiry_seconds"`
	RefreshTokenExpirySeconds int64  `mapstructure:"refresh_token_expiry_seconds"`
}

type Email struct {
	Sender      string `mapstructure:"sender"`
	AppPassword string `mapstructure:"app_password"`
	SmtpHost    string `mapstructure:"smtp_host"`
	SmtpPort    string `mapstructure:"smtp_port"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Firebase Firebase       `mapstructure:"firebase"`
	JWT      JWT            `mapstructure:"jwt"`
	Email    Email          `mapstructure:"email"`
	Redis    Redis          `mapstructure:"redis"`
}

func LoadConfig(path string) *Config {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Failed to read config file")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Err(err).Msg("Failed to unmarshal config")
	}

	return &cfg
}

package config

import (
	"fmt"
	"os"

	"github.com/misalima/nano-link-backend/src/infra/logger"
)

type PostgresConfig struct {
	User     string
	Host     string
	Password string
	Port     string
	DBName   string
}

type ServerConfig struct {
	Host string
	Port string
}

type LoggerConfig struct {
	Environment string
}

type Config struct {
	PostgresConfig PostgresConfig
	ServerConfig   ServerConfig
	LoggerConfig   LoggerConfig
}

func LoadConfig() *Config {
	postgresCfg := PostgresConfig{
		User:     getEnvOrDefault("PG_USER", "postgres"),
		Host:     getEnvOrDefault("PG_HOST", "localhost"),
		Password: getEnvOrDefault("PG_PASSWORD", "postgres"),
		Port:     getEnvOrDefault("PG_PORT", "5432"),
		DBName:   getEnvOrDefault("PG_NAME", "nano"),
	}

	serverCfg := ServerConfig{
		Host: getEnvOrDefault("SERVER_HOST", "localhost"),
		Port: getEnvOrDefault("SERVER_PORT", "8080"),
	}

	loggerCfg := LoggerConfig{
		Environment: getEnvOrDefault("ENVIRONMENT", "development"),
	}

	return &Config{
		PostgresConfig: postgresCfg,
		ServerConfig:   serverCfg,
		LoggerConfig:   loggerCfg,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		logger.Infof("Couldn't find environment value for: %s. Using default value: %s", key, defaultValue)
		return defaultValue
	}
	return value
}

func (cfg *Config) GetConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PostgresConfig.User,
		cfg.PostgresConfig.Password,
		cfg.PostgresConfig.Host,
		cfg.PostgresConfig.Port,
		cfg.PostgresConfig.DBName,
	)
}

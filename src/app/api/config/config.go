package config

import (
	"fmt"
	"log"
	"os"
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

type Config struct {
	PostgresConfig PostgresConfig
	ServerConfig   ServerConfig
}

func LoadConfig() *Config {
	postgresCfg := PostgresConfig{
		User:     getEnvOrDefault("DB_USER", "postgres"),
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
		Port:     getEnvOrDefault("DB_PORT", "5432"),
		DBName:   getEnvOrDefault("DB_NAME", "tasks_db"),
	}

	serverCfg := ServerConfig{
		Host: getEnvOrDefault("SERVER_HOST", "localhost"),
		Port: getEnvOrDefault("SERVER_PORT", "8080"),
	}

	return &Config{
		PostgresConfig: postgresCfg,
		ServerConfig:   serverCfg,
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Couldn't find environment value for: %s. Using default value: %s", value, defaultValue)
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

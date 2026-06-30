package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Auth     AuthConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Host string
	Port string
}
type AuthConfig struct {
	JWTSecret            string
	JWTExpirationMinutes int
}

func Load() (Config, error) {
	cfg := Config{
		Server: ServerConfig{
			Port: envOrDefault("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     envOrDefault("DATABASE_HOST", "localhost"),
			Port:     envOrDefault("DATABASE_PORT", "5432"),
			User:     envOrDefault("DATABASE_USER", "taskapi"),
			Password: envOrDefault("DATABASE_PASSWORD", "taskapi"),
			Name:     envOrDefault("DATABASE_NAME", "taskapi"),
		},
		Redis: RedisConfig{
			Host: envOrDefault("REDIS_HOST", "localhost"),
			Port: envOrDefault("REDIS_PORT", "6379"),
		},
		Auth: AuthConfig{
			JWTSecret:            envOrDefault("JWT_SECRET", "dev-jwt-secret"),
			JWTExpirationMinutes: envIntOrDefault("JWT_EXPIRATION_MINUTES", 60),
		},
	}

	return cfg, nil
}

func envOrDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return defaultValue
	}
	return value
}

func envIntOrDefault(key string, defaultValue int) int {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return defaultValue
	}
	return parsed
}

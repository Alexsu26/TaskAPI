package config

import "os"

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
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

func Load() (Config, error) {
	cfg := Config{
		Server: ServerConfig{
			Port: envOrDefault("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     envOrDefault("DATABASE_HOST", "localhost"),
			Port:     envOrDefault("DATABASE_PORT", "5432"),
			User:     envOrDefault("DATABASE_USER", "root"),
			Password: envOrDefault("DATABASE_PASSWORD", ""),
			Name:     envOrDefault("DATABASE_NAME", ""),
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

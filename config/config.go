package config

import (
	"os"
)

type Config struct {
	ServerAddress string
}

func LoadConfig() Config {
	return Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

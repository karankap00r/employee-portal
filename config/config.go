package config

import (
	"os"
)

type Config struct {
	ServerAddress string
}

func LoadConfig() Config {
	return Config{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}
}

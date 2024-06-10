package config

import (
	"os"
	"sync"
)

type Config struct {
	ServerAddress string
}

var (
	_    Config
	once sync.Once
)

func init() {
	LoadConfig()
}

func LoadConfig() {
	once.Do(func() {
		_ = Config{
			ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		}
	})
}

/****************************************************
	Utility Methods
*****************************************************/

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

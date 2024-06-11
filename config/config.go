package config

import (
	"os"
)

type Config struct {
	ServerAddress string
}

var (
	config Config
)

func init() {
	loadConfig()
}

func GetConfig() Config {
	return config
}

func loadConfig() {
	//todo: load config from config file
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

package config

import (
	"os"
)

// Config represents the configuration of the application
type Config struct {
	ServerAddress string
}

// Config instance
var (
	config Config
)

// init initializes the config
func init() {
	loadConfig()
}

// GetConfig returns the config instance
func GetConfig() Config {
	return config
}

// loadConfig loads the configuration
func loadConfig() {
	//todo: load config from config file
}

/****************************************************
	Utility Methods
*****************************************************/

// getEnv returns the value of the environment variable with the given key
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

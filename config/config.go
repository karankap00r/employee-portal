package config

import (
	"database/sql"
	"log"
	"os"
	"sync"
)

type Config struct {
	ServerAddress string
}

var config Config
var once sync.Once

var db *sql.DB

func init() {
	LoadConfig()
	initDatabase()
}

// Initialize SQLite database
func initDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatal(err)
	}
}

func LoadConfig() {
	once.Do(func() {
		config = Config{
			ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		}
	})
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

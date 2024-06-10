package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os/exec"
	"sync"
)

// databaseConnection is the database connection
var databaseConnection *sql.DB

// dbOnce is used to ensure that the database is initialized only once
var dbOnce sync.Once

// InitDB initializes the database and runs migrations
func InitDB() {
	var err error

	dbOnce.Do(func() {
		databaseConnection, err = sql.Open("sqlite3", "./employees.db")
		if err != nil {
			log.Fatalf("Failed to open database: %v", err)
		}
	})

	migrateDB(err)
}

// migrateDB runs Flyway migrations
func migrateDB(err error) {
	// Run Flyway migrations
	cmd := exec.Command("flyway", "migrate")
	cmd.Stdout = log.Writer()
	cmd.Stderr = log.Writer()
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run Flyway migrations: %v", err)
	}
}

// CloseDB closes the database connection
func CloseDB() {
	if err := databaseConnection.Close(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return databaseConnection
}

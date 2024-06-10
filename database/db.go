package database

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
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

	migrateDB()
}

// migrateDB runs Flyway migrations
func migrateDB() {
	driver, err := sqlite3.WithInstance(databaseConnection, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://github.com/karankap00r/employee_portal/migrations",
		"sqlite3", driver)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}

// CloseDB closes the database connection
func CloseDB() {
	if err := databaseConnection.Close(); err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}

	// Delete the database file
	if err := os.Remove("./employees.db"); err != nil {
		log.Fatalf("Failed to delete database file: %v", err)
	}
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return databaseConnection
}

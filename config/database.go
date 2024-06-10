package config

import (
	"database/sql"
	"sync"
)

var db *sql.DB
var dbOnce sync.Once

// InitializeDB sets up the database connection
func InitializeDB(database *sql.DB) {
	dbOnce.Do(func() {
		db = database
	})
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return db
}

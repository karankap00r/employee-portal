package main

import (
	"github.com/karankap00r/employee_portal/database"
	"github.com/karankap00r/employee_portal/server"

	"os"
	"os/signal"
	"syscall"

	_ "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/swaggo/http-swagger"
)

func main() {

	// Initialize the database
	database.InitDB()
	defer database.CloseDB()

	// Start server to initialise REST endpoints
	go server.Start()

	// Wait for interrupt signal to gracefully shutdown the server.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}

package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/config"

	//"github.com/karankap00r/employee_portal/server"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/api"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/swaggo/http-swagger"
)

// @title Employee API
// @version 1.0
// @description This is a sample server for managing employees.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {

	// Initialize SQLite database
	db, err := sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	createTable := `CREATE TABLE IF NOT EXISTS employees (
        id TEXT PRIMARY KEY,
        name TEXT,
        age INTEGER,
        dept TEXT
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	config.InitializeDB(db)

	//cfg := config.LoadConfig()

	//go server.Start(cfg)

	go func() {
		r := mux.NewRouter()

		r.HandleFunc("/employees", api.GetEmployees).Methods(http.MethodGet)
		r.HandleFunc("/employees/{id}", api.GetEmployee).Methods(http.MethodGet)
		r.HandleFunc("/employees", api.CreateEmployee).Methods(http.MethodPost)
		r.HandleFunc("/employees/{id}", api.UpdateEmployee).Methods(http.MethodPut)

		r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

		log.Println("Server started at port 8000")

		// Start server
		log.Fatal(http.ListenAndServe(":8000", r))
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}

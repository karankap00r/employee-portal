package main

import (
	"github.com/gorilla/mux"
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

	// Mock data
	api.Employees = append(api.Employees, api.Employee{ID: "1", Name: "John Doe", Age: 30, Dept: "Engineering"})
	api.Employees = append(api.Employees, api.Employee{ID: "2", Name: "Jane Doe", Age: 25, Dept: "Marketing"})

	//cfg := config.LoadConfig()

	//go server.Start(cfg)

	go func() {
		r := mux.NewRouter()

		r.HandleFunc("/employees", api.GetEmployees).Methods(http.MethodGet)
		r.HandleFunc("/employees/{id}", api.GetEmployee).Methods(http.MethodGet)
		r.HandleFunc("/employees", api.CreateEmployee).Methods(http.MethodPost)
		r.HandleFunc("/employees/{id}", api.UpdateEmployee).Methods(http.MethodPut)

		r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

		// Serve the docs folder
		fs := http.FileServer(http.Dir("./docs"))
		r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

		log.Println("Server started at port 8000")

		// Start server
		log.Fatal(http.ListenAndServe(":8000", r))
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-interrupt
}

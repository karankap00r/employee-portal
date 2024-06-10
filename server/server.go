package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/api"
	"github.com/karankap00r/employee_portal/config"
)

func Start(config config.Config) {
	// Initialize the router
	r := mux.NewRouter()
	// Route handlers & endpoints
	r.HandleFunc("/employees", api.GetEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", api.GetEmployee).Methods("GET")
	r.HandleFunc("/employees", api.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", api.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", api.DeleteEmployee).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

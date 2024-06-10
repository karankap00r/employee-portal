package server

import (
	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/api"
	"github.com/karankap00r/employee_portal/config"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

func Start(_ config.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/employees", api.GetEmployees).Methods(http.MethodGet)
	r.HandleFunc("/employees/{id}", api.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employees", api.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employees/{id}", api.UpdateEmployee).Methods(http.MethodPut)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started at port 8000")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

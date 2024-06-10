package server

import (
	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/api"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
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

func Start() {
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

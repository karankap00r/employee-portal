package server

import (
	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/api"
	"github.com/karankap00r/employee_portal/database"
	service "github.com/karankap00r/employee_portal/service/employee"
	"github.com/karankap00r/employee_portal/storage/repository"
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

	// Initialize repository, service, and handlers
	employeeRepo := repository.NewEmployeeRepository(database.GetDB())
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeHandler := api.NewEmployeeHandler(employeeService)

	r.HandleFunc("/employees", employeeHandler.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employees/{id}", employeeHandler.UpdateEmployee).Methods(http.MethodPut)
	r.HandleFunc("/employees/{id}", employeeHandler.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employees", employeeHandler.GetEmployees).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started at port 8000")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

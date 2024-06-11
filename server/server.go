package server

import (
	"github.com/gorilla/mux"
	"github.com/karankap00r/employee_portal/api"
	"github.com/karankap00r/employee_portal/database"
	"github.com/karankap00r/employee_portal/middleware"
	"github.com/karankap00r/employee_portal/service"
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
	dbConnection := database.GetDB()

	// Initialize repositories
	orgRepo := repository.NewOrgRepository(dbConnection)
	leaveRepo := repository.NewLeaveRepository(dbConnection)
	employeeRepo := repository.NewEmployeeRepository(dbConnection)
	remoteWorkRepo := repository.NewRemoteWorkRepository(dbConnection)

	// Initialize services
	leaveService := service.NewLeaveService(leaveRepo)
	employeeService := service.NewEmployeeService(employeeRepo)
	remoteWorkService := service.NewRemoteWorkService(remoteWorkRepo)

	// Initialize handlers
	leaveHandler := api.NewLeaveHandler(leaveService)
	employeeHandler := api.NewEmployeeHandler(employeeService)
	remoteWorkHandler := api.NewRemoteWorkHandler(remoteWorkService)

	r := mux.NewRouter()

	r.Use(middleware.OrgResolver(orgRepo))

	r.HandleFunc("/employee", employeeHandler.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee/{employeeID}", employeeHandler.UpdateEmployee).Methods(http.MethodPut)
	r.HandleFunc("/employee/{employeeID}", employeeHandler.GetEmployee).Methods(http.MethodGet)
	r.HandleFunc("/employees", employeeHandler.GetAllEmployees).Methods(http.MethodGet)

	r.HandleFunc("/leave-balance", leaveHandler.GetLeaveBalance).Methods(http.MethodGet)
	r.HandleFunc("/leave-request", leaveHandler.RaiseLeaveRequest).Methods(http.MethodPost)
	r.HandleFunc("/leave/{action}", leaveHandler.UpdateLeaveRequest).Methods(http.MethodPut)
	r.HandleFunc("/leaves-requests", leaveHandler.GetLeavesInRange).Methods(http.MethodGet)

	r.HandleFunc("/remote-work-balance", remoteWorkHandler.GetRemoteWorkBalance).Methods(http.MethodGet)
	r.HandleFunc("/remote-work-request", remoteWorkHandler.RaiseRemoteWorkRequest).Methods(http.MethodPost)
	r.HandleFunc("/remote-work/{action}", remoteWorkHandler.UpdateRemoteWorkRequest).Methods(http.MethodPut)
	r.HandleFunc("/remote-work-requests", remoteWorkHandler.GetRemoteWorkRequestsInRange).Methods(http.MethodGet)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server started at port 8000")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
